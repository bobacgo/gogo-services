package network_test

import (
	"testing"

	"github.com/gogoclouds/gogo-services/common-lib/pkg/network"
)

func TestGetOutBoundIP(t *testing.T) {
	ip, err := network.OutBoundIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}
