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
	cache      redis.Cmdable
	key        string        // 错误密码存放的key
	expiration time.Duration // 限制时长(在有效的错误次数范围内,每次错误都会刷新)
	limit      int64         // 错误次数限制
	errCount   int64         // 尝试次数
	OnErr      func(error)
}

func NewPasswdVerifier(rdb redis.Cmdable, limit int64) *PasswdVerifier {
	return &PasswdVerifier{cache: rdb, limit: limit, expiration: 24 * time.Hour}
}

// BcryptVerifyWithCount 验证密码统计错误次数
func (h *PasswdVerifier) BcryptVerifyWithCount(ctx context.Context, hash, password string) bool {
	if len(hash) <= 8 {
		h.fail(ctx)
		return false
	}
	hash, salt := h.parsePwd(hash)
	if !util.BcryptVerify(salt, hash, password) {
		h.fail(ctx)
		return false
	}
	h.delIncr(ctx)
	return true
}

// BcryptVerify 验证密码
func (h *PasswdVerifier) BcryptVerify(hash, password string) bool {
	hash, salt := h.parsePwd(hash)
	return util.BcryptVerify(salt, hash, password)
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
	h.errCount, err = h.cache.Incr(ctx, h.key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		h.OnErr(err)
		return
	}
	if err = h.cache.Expire(ctx, h.key, h.expiration).Err(); err != nil && !errors.Is(err, redis.Nil) {
		h.OnErr(err)
		return
	}
	if h.errCount >= h.limit {
		h.OnErr(ErrPasswdLimit)
	}
}

func (h *PasswdVerifier) delIncr(ctx context.Context) {
	if err := h.cache.Del(ctx, h.key).Err(); err != nil && !errors.Is(err, redis.Nil) {
		h.OnErr(err)
	}
}

func (h *PasswdVerifier) parsePwd(hashPwd string) (hash, salt string) {
	if len(hashPwd) <= 8 {
		return
	}
	return hashPwd[:len(hashPwd)-8], hashPwd[len(hashPwd)-8:]
}
