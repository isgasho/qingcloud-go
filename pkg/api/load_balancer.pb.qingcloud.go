// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: load_balancer.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type LoadBalancerServiceInterface interface {
	CreateLoadBalancer(in *CreateLoadBalancerInput) (out *CreateLoadBalancerOutput, err error)
	DescribeLoadBalancers(in *DescribeLoadBalancersInput) (out *DescribeLoadBalancersOutput, err error)
	DeleteLoadBalancers(in *DeleteLoadBalancersInput) (out *DeleteLoadBalancersOutput, err error)
	ModifyLoadBalancerAttributes(in *ModifyLoadBalancerAttributesInput) (out *ModifyLoadBalancerAttributesOutput, err error)
	StartLoadBalancers(in *StartLoadBalancersInput) (out *StartLoadBalancersOutput, err error)
	StopLoadBalancers(in *StopLoadBalancersInput) (out *StopLoadBalancersOutput, err error)
	UpdateLoadBalancers(in *UpdateLoadBalancersInput) (out *UpdateLoadBalancersOutput, err error)
	ResizeLoadBalancers(in *ResizeLoadBalancersInput) (out *ResizeLoadBalancersOutput, err error)
	AssociateEipsToLoadBalancer(in *AssociateEipsToLoadBalancerInput) (out *AssociateEipsToLoadBalancerOutput, err error)
	DissociateEipsFromLoadBalancer(in *DissociateEipsFromLoadBalancerInput) (out *DissociateEipsFromLoadBalancerOutput, err error)
	AddLoadBalancerListeners(in *AddLoadBalancerListenersInput) (out *AddLoadBalancerListenersOutput, err error)
	DescribeLoadBalancerListeners(in *DescribeLoadBalancerListenersInput) (out *DescribeLoadBalancerListenersOutput, err error)
	DeleteLoadBalancerListeners(in *DeleteLoadBalancerListenersInput) (out *DeleteLoadBalancerListenersOutput, err error)
	ModifyLoadBalancerListenerAttributes(in *ModifyLoadBalancerListenerAttributesInput) (out *ModifyLoadBalancerListenerAttributesOutput, err error)
	AddLoadBalancerBackends(in *AddLoadBalancerBackendsInput) (out *AddLoadBalancerBackendsOutput, err error)
	DescribeLoadBalancerBackends(in *DescribeLoadBalancerBackendsInput) (out *DescribeLoadBalancerBackendsOutput, err error)
	DeleteLoadBalancerBackends(in *DeleteLoadBalancerBackendsInput) (out *DeleteLoadBalancerBackendsOutput, err error)
	ModifyLoadBalancerBackendAttributes(in *ModifyLoadBalancerBackendAttributesInput) (out *ModifyLoadBalancerBackendAttributesOutput, err error)
	CreateLoadBalancerPolicy(in *CreateLoadBalancerPolicyInput) (out *CreateLoadBalancerPolicyOutput, err error)
	DescribeLoadBalancerPolicies(in *DescribeLoadBalancerPoliciesInput) (out *DescribeLoadBalancerPoliciesOutput, err error)
	ModifyLoadBalancerPolicyAttributes(in *ModifyLoadBalancerPolicyAttributesInput) (out *ModifyLoadBalancerPolicyAttributesOutput, err error)
	ApplyLoadBalancerPolicy(in *ApplyLoadBalancerPolicyInput) (out *ApplyLoadBalancerPolicyOutput, err error)
	DeleteLoadBalancerPolicies(in *DeleteLoadBalancerPoliciesInput) (out *DeleteLoadBalancerPoliciesOutput, err error)
	AddLoadBalancerPolicyRules(in *AddLoadBalancerPolicyRulesInput) (out *AddLoadBalancerPolicyRulesOutput, err error)
	DescribeLoadBalancerPolicyRules(in *DescribeLoadBalancerPolicyRulesInput) (out *DescribeLoadBalancerPolicyRulesOutput, err error)
	ModifyLoadBalancerPolicyRuleAttributes(in *ModifyLoadBalancerPolicyRuleAttributesInput) (out *ModifyLoadBalancerPolicyRuleAttributesOutput, err error)
	DeleteLoadBalancerPolicyRules(in *DeleteLoadBalancerPolicyRulesInput) (out *DeleteLoadBalancerPolicyRulesOutput, err error)
	CreateServerCertificate(in *CreateServerCertificateInput) (out *CreateServerCertificateOutput, err error)
	DescribeServerCertificates(in *DescribeServerCertificatesInput) (out *DescribeServerCertificatesOutput, err error)
	ModifyServerCertificateAttributes(in *ModifyServerCertificateAttributesInput) (out *ModifyServerCertificateAttributesOutput, err error)
	DeleteServerCertificates(in *DeleteServerCertificatesInput) (out *DeleteServerCertificatesOutput, err error)
}

type LoadBalancerService struct {
	ServerInfo       *ServerInfo
	Properties       *LoadBalancerServiceProperties
	LastResponseBody string
}

func NewLoadBalancerService(server *ServerInfo, serviceProp *LoadBalancerServiceProperties) (p *LoadBalancerService) {
	return &LoadBalancerService{
		ServerInfo: server,
		Properties: serviceProp,
	}
}

func (p *LoadBalancerService) CreateLoadBalancer(input *CreateLoadBalancerInput) (output *CreateLoadBalancerOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateLoadBalancerOutput)

	err = client.CallMethod(nil, "CreateLoadBalancer", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DescribeLoadBalancers(input *DescribeLoadBalancersInput) (output *DescribeLoadBalancersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeLoadBalancersOutput)

	err = client.CallMethod(nil, "DescribeLoadBalancers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DeleteLoadBalancers(input *DeleteLoadBalancersInput) (output *DeleteLoadBalancersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteLoadBalancersOutput)

	err = client.CallMethod(nil, "DeleteLoadBalancers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ModifyLoadBalancerAttributes(input *ModifyLoadBalancerAttributesInput) (output *ModifyLoadBalancerAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyLoadBalancerAttributesOutput)

	err = client.CallMethod(nil, "ModifyLoadBalancerAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) StartLoadBalancers(input *StartLoadBalancersInput) (output *StartLoadBalancersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StartLoadBalancersOutput)

	err = client.CallMethod(nil, "StartLoadBalancers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) StopLoadBalancers(input *StopLoadBalancersInput) (output *StopLoadBalancersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StopLoadBalancersOutput)

	err = client.CallMethod(nil, "StopLoadBalancers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) UpdateLoadBalancers(input *UpdateLoadBalancersInput) (output *UpdateLoadBalancersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(UpdateLoadBalancersOutput)

	err = client.CallMethod(nil, "UpdateLoadBalancers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ResizeLoadBalancers(input *ResizeLoadBalancersInput) (output *ResizeLoadBalancersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ResizeLoadBalancersOutput)

	err = client.CallMethod(nil, "ResizeLoadBalancers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) AssociateEipsToLoadBalancer(input *AssociateEipsToLoadBalancerInput) (output *AssociateEipsToLoadBalancerOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AssociateEipsToLoadBalancerOutput)

	err = client.CallMethod(nil, "AssociateEipsToLoadBalancer", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DissociateEipsFromLoadBalancer(input *DissociateEipsFromLoadBalancerInput) (output *DissociateEipsFromLoadBalancerOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DissociateEipsFromLoadBalancerOutput)

	err = client.CallMethod(nil, "DissociateEipsFromLoadBalancer", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) AddLoadBalancerListeners(input *AddLoadBalancerListenersInput) (output *AddLoadBalancerListenersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AddLoadBalancerListenersOutput)

	err = client.CallMethod(nil, "AddLoadBalancerListeners", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DescribeLoadBalancerListeners(input *DescribeLoadBalancerListenersInput) (output *DescribeLoadBalancerListenersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeLoadBalancerListenersOutput)

	err = client.CallMethod(nil, "DescribeLoadBalancerListeners", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DeleteLoadBalancerListeners(input *DeleteLoadBalancerListenersInput) (output *DeleteLoadBalancerListenersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteLoadBalancerListenersOutput)

	err = client.CallMethod(nil, "DeleteLoadBalancerListeners", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ModifyLoadBalancerListenerAttributes(input *ModifyLoadBalancerListenerAttributesInput) (output *ModifyLoadBalancerListenerAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyLoadBalancerListenerAttributesOutput)

	err = client.CallMethod(nil, "ModifyLoadBalancerListenerAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) AddLoadBalancerBackends(input *AddLoadBalancerBackendsInput) (output *AddLoadBalancerBackendsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AddLoadBalancerBackendsOutput)

	err = client.CallMethod(nil, "AddLoadBalancerBackends", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DescribeLoadBalancerBackends(input *DescribeLoadBalancerBackendsInput) (output *DescribeLoadBalancerBackendsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeLoadBalancerBackendsOutput)

	err = client.CallMethod(nil, "DescribeLoadBalancerBackends", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DeleteLoadBalancerBackends(input *DeleteLoadBalancerBackendsInput) (output *DeleteLoadBalancerBackendsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteLoadBalancerBackendsOutput)

	err = client.CallMethod(nil, "DeleteLoadBalancerBackends", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ModifyLoadBalancerBackendAttributes(input *ModifyLoadBalancerBackendAttributesInput) (output *ModifyLoadBalancerBackendAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyLoadBalancerBackendAttributesOutput)

	err = client.CallMethod(nil, "ModifyLoadBalancerBackendAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) CreateLoadBalancerPolicy(input *CreateLoadBalancerPolicyInput) (output *CreateLoadBalancerPolicyOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateLoadBalancerPolicyOutput)

	err = client.CallMethod(nil, "CreateLoadBalancerPolicy", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DescribeLoadBalancerPolicies(input *DescribeLoadBalancerPoliciesInput) (output *DescribeLoadBalancerPoliciesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeLoadBalancerPoliciesOutput)

	err = client.CallMethod(nil, "DescribeLoadBalancerPolicies", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ModifyLoadBalancerPolicyAttributes(input *ModifyLoadBalancerPolicyAttributesInput) (output *ModifyLoadBalancerPolicyAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyLoadBalancerPolicyAttributesOutput)

	err = client.CallMethod(nil, "ModifyLoadBalancerPolicyAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ApplyLoadBalancerPolicy(input *ApplyLoadBalancerPolicyInput) (output *ApplyLoadBalancerPolicyOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ApplyLoadBalancerPolicyOutput)

	err = client.CallMethod(nil, "ApplyLoadBalancerPolicy", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DeleteLoadBalancerPolicies(input *DeleteLoadBalancerPoliciesInput) (output *DeleteLoadBalancerPoliciesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteLoadBalancerPoliciesOutput)

	err = client.CallMethod(nil, "DeleteLoadBalancerPolicies", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) AddLoadBalancerPolicyRules(input *AddLoadBalancerPolicyRulesInput) (output *AddLoadBalancerPolicyRulesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AddLoadBalancerPolicyRulesOutput)

	err = client.CallMethod(nil, "AddLoadBalancerPolicyRules", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DescribeLoadBalancerPolicyRules(input *DescribeLoadBalancerPolicyRulesInput) (output *DescribeLoadBalancerPolicyRulesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeLoadBalancerPolicyRulesOutput)

	err = client.CallMethod(nil, "DescribeLoadBalancerPolicyRules", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ModifyLoadBalancerPolicyRuleAttributes(input *ModifyLoadBalancerPolicyRuleAttributesInput) (output *ModifyLoadBalancerPolicyRuleAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyLoadBalancerPolicyRuleAttributesOutput)

	err = client.CallMethod(nil, "ModifyLoadBalancerPolicyRuleAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DeleteLoadBalancerPolicyRules(input *DeleteLoadBalancerPolicyRulesInput) (output *DeleteLoadBalancerPolicyRulesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteLoadBalancerPolicyRulesOutput)

	err = client.CallMethod(nil, "DeleteLoadBalancerPolicyRules", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) CreateServerCertificate(input *CreateServerCertificateInput) (output *CreateServerCertificateOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateServerCertificateOutput)

	err = client.CallMethod(nil, "CreateServerCertificate", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DescribeServerCertificates(input *DescribeServerCertificatesInput) (output *DescribeServerCertificatesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeServerCertificatesOutput)

	err = client.CallMethod(nil, "DescribeServerCertificates", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) ModifyServerCertificateAttributes(input *ModifyServerCertificateAttributesInput) (output *ModifyServerCertificateAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyServerCertificateAttributesOutput)

	err = client.CallMethod(nil, "ModifyServerCertificateAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *LoadBalancerService) DeleteServerCertificates(input *DeleteServerCertificatesInput) (output *DeleteServerCertificatesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteServerCertificatesOutput)

	err = client.CallMethod(nil, "DeleteServerCertificates", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
