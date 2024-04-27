package cache_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/gogoclouds/gogo-services/framework/app/cache"
)

func TestCache(t *testing.T) {
	che, err := cache.DefaultCache()

	if err != nil {
		t.Error(err)
	}
	if err = che.Set("foo", "bar", 3*time.Second); err != nil {
		t.Error(err)
	}
	if err = che.Set("foo1", "bar1", time.Second); err != nil {
		t.Error(err)
	}
	t.Log(che.Keys())

	var value string
	err = che.Get("foo1", &value)
	t.Log(value)

	var value1 string
	che.Del("foo")
	err = che.Get("foo", &value1)
	t.Log(value1)

	var value2 string
	time.Sleep(2 * time.Second)
	err = che.Get("foo1", &value2)
	t.Log(value2)

	che.Clear()
	t.Log(che.Keys())

	// --------------------------------

	m := map[string]any{
		"name": "wlj",
		"age":  23,
	}
	if err := che.Set("hash", m, 3*time.Second); err != nil {
		t.Error(err)
	}
	var hv map[string]any
	err = che.Get("hash", &hv)
	t.Log(hv)
}

// 解析出单位
func TestParseUnit(t *testing.T) {
	size := "512GB"
	re, _ := regexp.Compile("^[0-9]+")

	loc := re.FindStringIndex(size)
	// unit := string(re.ReplaceAll([]byte(size), []byte("")))
	t.Log(size[:loc[1]], size[loc[1]:])
}
