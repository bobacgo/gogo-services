package dns

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/polarismesh/polaris-go"
)

// register server

type providerServer struct {
	provider  polaris.ProviderAPI
	namespace string
	service   string
	host      string
	port      uint16
}

func (c providerServer) Register() {
	g.Log.Info("start to invoke register operation")
	registerRequest := new(polaris.InstanceRegisterRequest)
	registerRequest.Service = c.service
	registerRequest.Namespace = c.namespace
	registerRequest.Host = c.host
	registerRequest.Port = int(c.port)
	registerRequest.SetTTL(10)
	resp, err := c.provider.RegisterInstance(registerRequest)
	if err != nil {
		g.Log.Infof("fail to register instance, err is %v", err)
	}
	g.Log.Infof("register response: instanceId %s", resp.InstanceID)
}

func (c providerServer) RunMainLoop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, []os.Signal{
		syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGSEGV,
	}...)

	for s := range ch {
		g.Log.Infof("catch signal(%+v), stop servers", s)
		c.deregisterService()
		return
	}
}

func (c providerServer) deregisterService() {
	g.Log.Info("start to invoke deregister operation")
	deregisterRequest := new(polaris.InstanceDeRegisterRequest)
	deregisterRequest.Service = c.service
	deregisterRequest.Namespace = c.namespace
	deregisterRequest.Host = c.host
	deregisterRequest.Port = int(c.port)
	if err := c.provider.Deregister(deregisterRequest); err != nil {
		log.Fatalf("fail to deregister instance, err is %v", err)
	}
	g.Log.Info("deregister successfully.")
}
