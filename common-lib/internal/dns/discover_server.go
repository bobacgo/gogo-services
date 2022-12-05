package dns

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/polarismesh/polaris-go"
)

func (c DiscoverServer) Register() {
	g.Log.Info("start to invoke register operation")
	registerRequest := new(polaris.InstanceRegisterRequest)
	registerRequest.Service = c.Service
	registerRequest.Namespace = c.Namespace
	registerRequest.Host = c.Host
	registerRequest.Port = int(c.Port)
	registerRequest.SetTTL(10)
	resp, err := c.Provider.RegisterInstance(registerRequest)
	if err != nil {
		g.Log.Infof("fail to register instance, err is %v", err)
	}
	g.Log.Infof("register response: instanceId %s", resp.InstanceID)
}

func (c DiscoverServer) deregisterService() {
	g.Log.Info("start to invoke deregister operation")
	deregisterRequest := &polaris.InstanceDeRegisterRequest{}
	deregisterRequest.Service = c.Service
	deregisterRequest.Namespace = c.Namespace
	deregisterRequest.Host = c.Host
	deregisterRequest.Port = int(c.Port)
	if err := c.Provider.Deregister(deregisterRequest); err != nil {
		log.Fatalf("fail to deregister instance, err is %v", err)
	}
	g.Log.Info("deregister successfully.")
}

func (c DiscoverServer) RunMainLoop() {
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
