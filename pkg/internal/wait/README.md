# wait
`import "github.com/chai2010/qingcloud-go/pkg/wait"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/chai2010/qingcloud-go/pkg/api](./../api)
- [github.com/chai2010/qingcloud-go/pkg/status](./../status)

## <a name="pkg-index">Index</a>
* [func NewErrTimeout(err error) error](#NewErrTimeout)
* [func WaitForIntervalWorkDone(name string, timeout time.Duration, intervalWork func() (done bool, err error)) error](#WaitForIntervalWorkDone)
* [func WaitForOnetimeWork(name string, timeout time.Duration, onetimeWork func() error) (err error)](#WaitForOnetimeWork)
* [func WaitInstanceNetwork(server \*pb.ServerInfo, insId string, timeout time.Duration) error](#WaitInstanceNetwork)
* [func WaitInstanceStatus(server \*pb.ServerInfo, insId string, status statuspkg.InstanceStatus, timeout time.Duration) error](#WaitInstanceStatus)
* [func WaitJob(server \*pb.ServerInfo, jobId string, timeout time.Duration) error](#WaitJob)
* [func WaitLoadBalancerStatus(server \*pb.ServerInfo, lbId string, status statuspkg.LoadBalancerStatus, timeout time.Duration) error](#WaitLoadBalancerStatus)
* [type TimeoutError](#TimeoutError)

#### <a name="pkg-files">Package files</a>
[wait.go](./wait.go) [wait_ins.go](./wait_ins.go) [wait_job.go](./wait_job.go) [wait_lb.go](./wait_lb.go) 

## <a name="NewErrTimeout">func</a> [NewErrTimeout](./wait.go#L21)
``` go
func NewErrTimeout(err error) error
```

## <a name="WaitForIntervalWorkDone">func</a> [WaitForIntervalWorkDone](./wait.go#L33-L36)
``` go
func WaitForIntervalWorkDone(
    name string, timeout time.Duration,
    intervalWork func() (done bool, err error),
) error
```

## <a name="WaitForOnetimeWork">func</a> [WaitForOnetimeWork](./wait.go#L49-L52)
``` go
func WaitForOnetimeWork(
    name string, timeout time.Duration,
    onetimeWork func() error,
) (err error)
```

## <a name="WaitInstanceNetwork">func</a> [WaitInstanceNetwork](./wait_ins.go#L27-L30)
``` go
func WaitInstanceNetwork(
    server *pb.ServerInfo, insId string,
    timeout time.Duration,
) error
```

## <a name="WaitInstanceStatus">func</a> [WaitInstanceStatus](./wait_ins.go#L15-L19)
``` go
func WaitInstanceStatus(
    server *pb.ServerInfo, insId string,
    status statuspkg.InstanceStatus,
    timeout time.Duration,
) error
```

## <a name="WaitJob">func</a> [WaitJob](./wait_job.go#L15)
``` go
func WaitJob(server *pb.ServerInfo, jobId string, timeout time.Duration) error
```

## <a name="WaitLoadBalancerStatus">func</a> [WaitLoadBalancerStatus](./wait_lb.go#L15-L19)
``` go
func WaitLoadBalancerStatus(
    server *pb.ServerInfo, lbId string,
    status statuspkg.LoadBalancerStatus,
    timeout time.Duration,
) error
```

## <a name="TimeoutError">type</a> [TimeoutError](./wait.go#L12-L15)
``` go
type TimeoutError interface {
    Timeout() bool
    error
}
```

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)