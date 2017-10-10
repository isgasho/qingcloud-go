// +-------------------------------------------------------------------------
// | Copyright (C) 2017 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/user"
	"path"
	"time"

	"github.com/chai2010/qingcloud-go/config"
	qcservice "github.com/chai2010/qingcloud-go/service"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/mcnutils"
	"github.com/docker/machine/libmachine/ssh"
	"github.com/docker/machine/libmachine/state"
)

const (
	defaultImage = "xenialx64b"
	//defaultImage     = "trustysrvx64h"
	defaultZone         = "pek3a"
	defaultCPU          = 1
	defaultMemory       = 1024
	defaultOpTimeout    = 180 //second
	defaultVxNet        = "vxnet-0"
	defaultEIPBandwidth = 4 //MB
	dockerPort          = 2376
	swarmPort           = 3376
)

var defaultSecurityGroupRules = []*qcservice.SecurityGroupRule{
	{
		Priority: 0,
		Protocol: "icmp",
		Action:   "accept",
		Val1:     "8", //Echo
		Val2:     "0", //Echo request
		Val3:     "",
	},
	{
		Priority: 1,
		Protocol: "tcp",
		Action:   "accept",
		Val1:     "22",
		Val2:     "",
		Val3:     "",
	},
	{
		Priority: 2,
		Protocol: "tcp",
		Action:   "accept",
		Val1:     "2376",
		Val2:     "",
		Val3:     "",
	},
}

type Driver struct {
	*drivers.BaseDriver
	AccessKeyID     string
	SecretAccessKey string
	Zone            string
	Image           string
	CPU             int
	Memory          int
	LoginKeyPair    string
	VxNet           string
	InstanceID      string
	EIP             *qcservice.EIP
	SecurityGroup   *qcservice.SecurityGroup
	client          Client
}

type SSHKeyPair struct {
	ID string
}

// GetCreateFlags registers the flags this driver adds to
// "docker hosts create"

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	user, err := user.Current()
	if err != nil {
		log.Errorf("Get current user error: %s", err.Error())
	}
	defaultSSHKeyPath := path.Join(user.HomeDir, ".ssh/id_rsa")
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_ACCESS_KEY_ID",
			Name:   "qingcloud-access-key-id",
			Usage:  "QingCloud access key id",
		},
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_SECRET_ACCESS_KEY",
			Name:   "qingcloud-secret-access-key",
			Usage:  "QingCloud secret access key",
		},
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_ZONE",
			Name:   "qingcloud-zone",
			Usage:  "QingCloud zone",
			Value:  defaultZone,
		},
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_IMAGE",
			Name:   "qingcloud-image",
			Usage:  "Instance image ID",
			Value:  defaultImage,
		},
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_VXNET_ID",
			Name:   "qingcloud-vxnet-id",
			Usage:  "Vxnet id",
			Value:  defaultVxNet,
		},
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_LOGIN_KEYPAIR",
			Name:   "qingcloud-login-keypair",
			Usage:  "Login keypair id.",
		},
		mcnflag.StringFlag{
			EnvVar: "QINGCLOUD_SSH_KEYPATH",
			Name:   "qingcloud-ssh-keypath",
			Usage:  "SSH Key for Instance.",
			Value:  defaultSSHKeyPath,
		},
		mcnflag.IntFlag{
			Name:  "qingcloud-cpu",
			Usage: "QingCloud cpu count",
			Value: defaultCPU,
		},
		mcnflag.IntFlag{
			Name:  "qingcloud-memory",
			Usage: "QingCloud memory size in MB",
			Value: defaultMemory,
		},
	}
}

// SetConfigFromFlags configures the driver with the object that was returned
// by RegisterCreateFlags
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.AccessKeyID = flags.String("qingcloud-access-key-id")
	d.SecretAccessKey = flags.String("qingcloud-secret-access-key")
	d.Zone = flags.String("qingcloud-zone")
	d.VxNet = flags.String("qingcloud-vxnet-id")
	d.LoginKeyPair = flags.String("qingcloud-login-keypair")
	d.CPU = flags.Int("qingcloud-cpu")
	d.Memory = flags.Int("qingcloud-memory")
	d.SSHKeyPath = flags.String("qingcloud-ssh-keypath")
	d.Image = flags.String("qingcloud-image")
	d.SetSwarmConfigFromFlags(flags)
	return nil
}

func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		Image:  defaultImage,
		CPU:    defaultCPU,
		Memory: defaultMemory,
		Zone:   defaultZone,
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "qingcloud"
}

func (d *Driver) Config() *config.Config {
	config, err := config.New(d.AccessKeyID, d.SecretAccessKey)
	if err != nil {
		panic(fmt.Sprintf("init config error: %s", err.Error()))
	}
	return config
}

// PreCreateCheck allows for pre-create operations to make sure a driver is ready for creation
func (d *Driver) PreCreateCheck() error {
	if d.LoginKeyPair != "" {
		_, err := d.GetClient().DescribeKeyPair(d.LoginKeyPair)
		if err != nil {
			return err
		}
		if d.SSHKeyPath == "" {
			return errors.New("Param error: qingcloud-login-keypair param should work with qingcloud-ssh-keypath param.")
		}
	}
	if d.VxNet == "" {
		return errors.New("Param qingcloud-vxnet-id required.")
	}

	return nil
}

func (d *Driver) Create() error {
	log.Infof("Creating SSH key...")

	if d.LoginKeyPair == "" {
		err := d.createSSHKey()
		if err != nil {
			return err
		}
	}

	log.Infof("Creating QingCloud Instance...")

	client := d.GetClient()
	arg := &RunInstanceArg{
		CPU:          d.CPU,
		Memory:       d.Memory,
		ImageID:      d.Image,
		VxNet:        d.VxNet,
		LoginKeyPair: d.LoginKeyPair,
		InstanceName: d.MachineName,
	}
	ins, err := client.RunInstance(arg)
	if err != nil {
		return err
	}
	d.InstanceID = ins.InstanceId

	if d.VxNet == defaultVxNet {
		eip, err := client.BindEIP(d.InstanceID)
		if err != nil {
			return err
		}
		log.Infof("Bind EIP [%s] to Instance [%s]", eip.EipAddr, d.InstanceID)
		ins.Eip = eip
		d.EIP = eip
		sg, err := client.BindSecurityGroup(d.InstanceID, defaultSecurityGroupRules)
		if err != nil {
			return err
		}
		d.SecurityGroup = sg
		log.Infof("Bind SecurityGroup [%s] to Instance [%s]", sg.SecurityGroupId, d.InstanceID)
	}

	d.IPAddress = ins.Vxnets[0].PrivateIp
	if ins.Eip != nil && ins.Eip.EipAddr != "" {
		d.IPAddress = ins.Eip.EipAddr
	}
	d.MachineName = d.InstanceID

	log.Infof("Created Instance [%s] IPAddress: [%s]",
		d.InstanceID, d.IPAddress)
	d.checkOSEnv()

	return nil
}

func (d *Driver) checkOSEnv() error {
	log.Infof("Check OS Env on Instance [%s]", d.InstanceID)
	sshClient, err := drivers.GetSSHClientFromDriver(d)
	if err != nil {
		log.Errorf("Get ssh client for [%s] error: [%s]", d.InstanceID, err.Error())
		return err
	}
	// check access public network
	err = mcnutils.WaitForSpecific(func() bool {
		err := sshClient.Shell("ping -q -c 3 -W 10 get.docker.com")
		if err != nil {
			return false
		}
		return true
	}, (defaultOpTimeout / 10), 10*time.Second)
	if err != nil {
		log.Errorf("Ping get.docker.com on Instance [%s] error :[%s]", d.InstanceID, err.Error())
		return err
	}
	err = mcnutils.WaitForSpecific(func() bool {
		err := sshClient.Shell("apt-get update")
		if err != nil {
			//kill process for lock /var/lib/dpkg/lock
			sshClient.Shell("fuser -kw /var/lib/dpkg/lock")
			sshClient.Shell("fuser -kw /var/lib/apt/lists/lock")
			//dpkg interrupted, so reconfigure
			sshClient.Shell("dpkg --configure -a")
			sshClient.Shell("apt-get clean")
			return false
		}
		return true
	}, (defaultOpTimeout / 20), 20*time.Second)
	if err != nil {
		log.Errorf("Apt-get update on Instance [%s] error :[%s]", d.InstanceID, err.Error())
		return err
	}

	return nil
}

func (d *Driver) GetClient() Client {
	if d.client == nil {
		client, err := NewClient(d.Config(), d.Zone)
		if err != nil {
			panic(fmt.Sprintf("init client error: %s", err.Error()))
		}
		d.client = client
	}
	return d.client
}

func (d *Driver) getInstance() (*qcservice.Instance, error) {
	return d.GetClient().DescribeInstance(d.InstanceID)
}

func (d *Driver) createSSHKey() error {

	if d.SSHKeyPath == "" {
		log.Debugf("Creating New SSH Key")
		if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
			return err
		}
		d.SSHKeyPath = d.GetSSHKeyPath()
	} else {
		log.Debugf("Using SSHKeyPath: %s", d.SSHKeyPath)
		if d.LoginKeyPair != "" {
			log.Debugf("Using existing LoginKeyPair: %s", d.LoginKeyPair)
			return nil
		}
	}

	publicKey, err := ioutil.ReadFile(d.publicSSHKeyPath())
	if err != nil {
		return err
	}

	keyName := d.MachineName

	log.Debugf("Creating key pair: %s", keyName)
	keyPairID, err := d.GetClient().CreateKeyPair(keyName, string(publicKey))
	if err != nil {
		return err
	}
	d.LoginKeyPair = *keyPairID
	return nil
}

func (d *Driver) publicSSHKeyPath() string {
	return d.GetSSHKeyPath() + ".pub"
}

// GetURL returns a Docker compatible host URL for connecting to this host
// e.g. tcp://1.2.3.4:2376
func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	if ip == "" {
		return "", nil
	}
	return fmt.Sprintf("tcp://%s:%d", ip, dockerPort), nil
}

// GetState returns the state that the host is in (running, stopped, etc)
func (d *Driver) GetState() (state.State, error) {
	i, err := d.getInstance()
	if err != nil {
		return state.None, err
	}
	if i.Status == "" {
		return state.None, nil
	}
	switch i.Status {
	case INSTANCE_STATUS_PENDING:
		return state.Starting, nil
	case INSTANCE_STATUS_RUNNING:
		return state.Running, nil
	case INSTANCE_STATUS_STOPPED:
		return state.Stopped, nil
	case INSTANCE_STATUS_SUSPENDED, INSTANCE_STATUS_TERMINATED, INSTANCE_STATUS_CEASED:
		return state.Error, nil
	}
	return state.Error, nil
}

// Kill stops a host forcefully
func (d *Driver) Kill() error {
	return d.GetClient().StopInstance(d.InstanceID, true)
}

// Remove a host
func (d *Driver) Remove() error {
	err := d.GetClient().TerminateInstance(d.InstanceID)
	if err != nil {
		return err
	}
	if d.EIP != nil {
		err := d.GetClient().ReleaseEIP(d.EIP.EipId)
		if err != nil {
			log.Errorf("Release EIP [%+v] fail, err: [%s]", *d.EIP, err.Error())
		}
	}
	if d.SecurityGroup != nil {
		err := d.GetClient().DeleteSecurityGroup(d.SecurityGroup.SecurityGroupId)
		if err != nil {
			log.Errorf("Delete SecurityGroup [%+v] fail, err: [%s]", *d.SecurityGroup, err.Error())
		}
	}
	return nil
}

// Restart a host. This may just call Stop(); Start() if the provider does not
// have any special restart behaviour.
func (d *Driver) Restart() error {
	return d.GetClient().RestartInstance(d.InstanceID)
}

// Start a host
func (d *Driver) Start() error {
	return d.GetClient().StartInstance(d.InstanceID)
}

// Stop a host gracefully
func (d *Driver) Stop() error {
	return d.GetClient().StopInstance(d.InstanceID, false)
}
