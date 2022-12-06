package pkg_test

import (
	"testing"

	"github.com/gogoclouds/gogo-services/common-lib/pkg"
)

func TestGetOutBoundIP(t *testing.T) {
	ip, err := pkg.Addr.GetOutBoundIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}

func TestParseAddr(t *testing.T) {
	ip, port := pkg.Addr.Parse("127.0.0.1:8080")
	if ip != "127.0.0.1" || port != 8080 {
		t.Fail()
	}
	t.Logf("ip: %s, port: %d", ip, port)
}
