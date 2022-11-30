package stream_test

import (
	"github.com/gogoclouds/gogo-services/common-lib/pkg/stream"
	"testing"
)

func TestStream(t *testing.T) {
	stream.Of("mysql", "redis", "kafka", "go", "go").
		Distinct().
		Filter(func(s string) bool {
			return s != "go"
		}).
		Each(func(i int, v string) {
			t.Log(i, v)
		})
	//t.Log(list)
}
