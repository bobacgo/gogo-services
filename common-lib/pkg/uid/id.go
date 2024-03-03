package uid

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sony/sonyflake"
)

// RandSeqID
// 生成随机序列（包括：大小写字母、数字）
func RandSeqID(n int) func() string {
	letters := []rune("0123456789abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPGRSTUVWXYZ")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[r.Intn(len(letters))]
		}
		return string(b)
	}
}

// UUID 不包含"-"
// 小写字母 + 数字
func UUID() string {
	newId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return newId
}

func Snowflake(startTime time.Time, machineID uint16) *sonyflake.Sonyflake {
	st := sonyflake.Settings{
		//StartTime: startTime,
		//MachineID: func() (uint16, error) {
		//	return machineID, nil
		//},
	}
	sf := sonyflake.NewSonyflake(st)
	return sf
}

// RandNumber 指定范围的随机数
func RandNumber(start, end int) func() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() int {
		if start >= end {
			return end
		}
		num := r.Intn(end-start) + start
		return num
	}
}
