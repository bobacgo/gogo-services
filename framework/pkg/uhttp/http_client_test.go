package uhttp_test

import (
	"testing"

	"github.com/gogoclouds/gogo-services/framework/pkg/uhttp"
)

func TestHttp(t *testing.T) {
	data, err := uhttp.Get[any]("https://www.baidu.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
