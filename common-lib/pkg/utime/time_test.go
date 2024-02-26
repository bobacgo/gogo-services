package utime_test

import (
	"fmt"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/utime"
	"testing"
	"time"
)

func TestZeroHour(t *testing.T) {
	remain := utime.ZeroHour(1).Unix() - time.Now().Unix()
	fmt.Println(remain)
}
