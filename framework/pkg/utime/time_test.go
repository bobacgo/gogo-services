package utime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogoclouds/gogo-services/framework/pkg/utime"
)

func TestZeroHour(t *testing.T) {
	remain := utime.ZeroHour(1).Unix() - time.Now().Unix()
	fmt.Println(remain)
}
