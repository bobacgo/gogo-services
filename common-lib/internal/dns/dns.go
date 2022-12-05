package dns

import "github.com/polarismesh/polaris-go"

// ConfigServer 配置中心
type ConfigServer struct{}

// DiscoverServer 注册中心
type DiscoverServer struct {
	Provider  polaris.ProviderAPI
	Namespace string
	Service   string
	Host      string
	Port      uint16
}
