package pkg_test

import (
	"testing"

	"github.com/gogoclouds/gogo-services/common-lib/pkg"
)

type person struct {
	Name string
	Age  uint8
}

func TestDistinct(t *testing.T) {
	arr := []string{"mysql", "redis", "mysql"}
	sd := pkg.StreamSlice(arr).
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
	pd := pkg.StreamSlice(p).Distinct().
		Peek(func(o *person) {
			o.Age = 19
		}).List()
	t.Log(pd) // [{fei.zhang 18} {bei.liu 22}]
}