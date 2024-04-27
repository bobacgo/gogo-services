package dns

import (
	"github.com/gogoclouds/gogo-services/framework/internal/dns/discover"
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/api"
)

type Server struct {
	Ctx api.SDKContext
}

// ConfigServer 配置中心
func (dns Server) Config() *configServer {
	return &configServer{ctx: dns.Ctx}
}

// DiscoverServer 注册中心 - 注册服务
func (dns Server) Provider(namespace, service, ip string, port uint16) *providerServer {
	pa := polaris.NewProviderAPIByContext(dns.Ctx)
	return &providerServer{
		provider:  pa,
		namespace: namespace,
		service:   service,
		host:      ip,
		port:      port,
	}
}

// DiscoverServer 注册中心 - 发现服务
func (dns Server) Consumer(namespace string) *discover.ConsumerServer {
	ca := polaris.NewConsumerAPIByContext(dns.Ctx)
	return &discover.ConsumerServer{
		Consumer:  ca,
		Namespace: namespace,
	}
}
