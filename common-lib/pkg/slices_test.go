package pkg_test

import (
	"github.com/gogoclouds/gogo-services/common-lib/pkg"
	"testing"
)

type person struct {
	Name string
	Age  uint8
}

func TestDistinct(t *testing.T) {
	arr := []string{"mysql", "redis", "mysql"}
	sd := pkg.Slices[string]{}.Distinct(arr)
	t.Log(sd) // [mysql redis]

	p := []person{
		{"fei.zhang", 18},
		{"fei.zhang", 18},
		{"bei.liu", 22},
	}
	pd := pkg.Slices[person]{}.Distinct(p)
	t.Log(pd) // [{fei.zhang 18} {bei.liu 22}]
}