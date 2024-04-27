package cache

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// https://zhuanlan.zhihu.com/p/635603181

type Cache interface {
	// SetMaxMemory size : 1KB 100KB 1M 2MB 1GB
	SetMaxMemory(size string) bool

	Set(key string, val any, expire time.Duration) error

	Get(key string, result any) error

	Del(key string) bool

	Exists(key string) bool

	Clear() bool
	// Keys 获取所有缓存中 key 的数量
	Keys() int64
}

var defaultSize = "512M"

func DefaultCache() (Cache, error) {
	return NewFreeCache(defaultSize)
}

func parseUnit(memSize string) (int, error) {
	re, _ := regexp.Compile("[0-9]+")
	loc := re.FindStringIndex(memSize)
	if len(loc) != 2 {
		return 0, fmt.Errorf("unit parse not exist: %s len(loc) = %d", memSize, len(loc))
	}
	num, _ := strconv.Atoi(memSize[:loc[1]])
	unit := strings.ToUpper(memSize[loc[1]:])
	switch unit {
	case "B":
		return num, nil
	case "KB", "K":
		return num << 10, nil
	case "MB", "M":
		return num << 20, nil
	case "GB", "G":
		return num << 30, nil
	case "TB", "T":
		return num << 40, nil
	case "PB", "P":
		return num << 50, nil
	default:
		return 0, fmt.Errorf("unit parse not exist: %s -> %s", memSize, memSize[loc[1]:])
	}
}
