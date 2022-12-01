package stream_test

import (
	"github.com/gogoclouds/gogo-services/common-lib/pkg/stream"
	"testing"
)

type person struct {
	Name string
	Age  uint8
}

func TestStream(t *testing.T) {
	stream.Of("mysql", "redis", "kafka", "go", "go", "java").
		Distinct().
		Filter(func(s string) bool {
			return s != "java"
		}).
		Reverse().
		Limit(10).
		Each(func(i int, v string) {
			t.Log(i, v)
		})
	//t.Log(list)
}

func TestDistinct(t *testing.T) {
	arr := []string{"mysql", "redis", "mysql"}
	sd := stream.New(arr).
		Filter(func(str string) bool {
			return str != "redis"
		}).
		Distinct().
		List()
	t.Log(sd) // [mysql redis]

	p := []person{
		{"fei.zhang", 18},
		{"fei.zhang", 18},
		{"bei.liu", 22},
	}
	pd := stream.Of(p...).Distinct().List()
	t.Log(pd) // [{fei.zhang 18} {bei.liu 22}]
}