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
	sd := pkg.StreamSlice(arr).Distinct().List()
	t.Log(sd) // [mysql redis]

	p := []person{
		{"fei.zhang", 18},
		{"fei.zhang", 18},
		{"bei.liu", 22},
	}
	pd := pkg.StreamSlice(p).Distinct().List()
	t.Log(pd) // [{fei.zhang 18} {bei.liu 22}]
}
