package security

import (
	"context"
	"errors"
	"github.com/gogoclouds/gogo/pkg/util"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrPasswdLimit = errors.New("password error limit")

// PasswdVerifier 登录密码验证器
// 1.对密码进行hash加密
// 2.随机生成盐
// 3.密码错误次数限制(依赖Redis)
type PasswdVerifier struct {
	rdb        redis.Cmdable
	key        string        // 错误密码存放的key
	expiration time.Duration // 限制时长(在有效的错误次数范围内,每次错误都会刷新)
	limit      int64         // 错误次数限制
	errCount   int64         // 尝试次数
	OnErr      func(error)
}

func NewPasswdVerifier(rdb redis.Cmdable, limit int64) *PasswdVerifier {
	return &PasswdVerifier{rdb: rdb, limit: limit, expiration: 24 * time.Hour}
}

// BcryptVerify 验证密码
func (h *PasswdVerifier) BcryptVerify(ctx context.Context, hash, password string) bool {
	var err error
	h.errCount, err = h.rdb.Get(ctx, h.key).Int64()
	if err != nil {
		h.OnErr(err)
		return false
	}
	if h.errCount > h.limit {
		h.OnErr(ErrPasswdLimit)
		return false
	}
	if len(password) < 8 {
		return false
	}
	if !util.BcryptVerify(hash[len(password)-8:], hash[:len(password)-8], password) {
		h.fail(ctx)
		return false
	}
	h.delIncr(ctx)
	return true
}

// BcryptHash 密码加密
func (h *PasswdVerifier) BcryptHash(passwd string) string {
	hash, salt := util.BcryptHash(passwd)
	return hash + salt
}

// GetErrCount 获取密码错误的次数
func (h *PasswdVerifier) GetErrCount() int64 {
	return h.errCount
}

// GetRemainCount 获取密码剩余的错误次数
func (h *PasswdVerifier) GetRemainCount() int64 {
	remainCount := h.limit - h.errCount
	if remainCount < 0 {
		remainCount = 0
	}
	return remainCount
}

// SetKey 设置key
func (h *PasswdVerifier) SetKey(key string, expiration time.Duration) {
	h.key = key
	h.expiration = expiration
}

func (h *PasswdVerifier) fail(ctx context.Context) {
	var err error
	h.errCount, err = h.rdb.Incr(ctx, h.key).Result()
	if err != nil {
		h.OnErr(err)
		return
	}
	if err = h.rdb.Expire(ctx, h.key, h.expiration).Err(); err != nil {
		h.OnErr(err)
		return
	}
	if h.errCount > h.limit {
		h.OnErr(ErrPasswdLimit)
	}
}

func (h *PasswdVerifier) delIncr(ctx context.Context) {
	if err := h.rdb.Del(ctx, h.key).Err(); err != nil {
		h.OnErr(err)
	}
}
