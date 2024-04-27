package discover

import (
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
)

// get server

type ConsumerServer struct {
	Consumer  polaris.ConsumerAPI
	Namespace string
	service   string
}

// GetOneInstance
// 每次仅获取一个可用服务提供者实例，该方法会依次执行路由、负载均衡流程。
// 该方法默认会过滤掉不健康、隔离、权重为0、被熔断的实例。
func (c ConsumerServer) GetOneInstance(service string) (model.Instance, error) {
	retryCount := 3
	gair := new(polaris.GetOneInstanceRequest)
	gair.Service = service
	gair.Namespace = c.Namespace
	gair.RetryCount = &retryCount
	instance, err := c.Consumer.GetOneInstance(gair)
	if err != nil {
		return nil, err
	}
	return instance.GetInstance(), err
}