package uhttp_test

import (
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uhttp"
	"testing"
)

func TestHttp(t *testing.T) {
	data, err := uhttp.Get[any]("https://www.baidu.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
