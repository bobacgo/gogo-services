package main

import (
	"github.com/gogoclouds/gogo-services/common-lib/dns"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/mapset"
)

var (
	dnsConfigFilePath = "../configs/polaris.yaml"
	dnsNamespace      = "default"
	fileGroup         = "gogo_v1.0.0"
	configFilenames   = []string{
		"admin-service.yaml", "common.yaml", "mysql.yaml", "redis.yaml", "test.yaml",
	}
)

func main() {
	server := dns.Server
	fileMetadata := &dns.FileMetadata{
		Namespace: dnsNamespace, FileGroup: fileGroup, FileNameSet: mapset.Of(configFilenames...),
	}
	server.LoadConfig(dnsConfigFilePath, fileMetadata)
	select {}
}
