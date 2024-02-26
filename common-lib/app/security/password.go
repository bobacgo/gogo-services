package security

import (
	"context"
	"errors"
	"github.com/gogoclouds/gogo/pkg/util"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrPasswdLimit = errors.New("password error limit")

type PasswdHelper struct {
	rdb        redis.UniversalClient
	key        string
	expiration time.Duration
	limit      int64
	errCount   int64
	OnErr      func(error)
}

func NewPasswdHelper(rdb redis.UniversalClient, limit int64) *PasswdHelper {
	return &PasswdHelper{rdb: rdb, limit: limit, expiration: 24 * time.Hour}
}

func (h *PasswdHelper) BcryptVerify(ctx context.Context, hash, password string) bool {
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

func (h *PasswdHelper) BcryptHash(passwd string) string {
	hash, salt := util.BcryptHash(passwd)
	return hash + salt
}

func (h *PasswdHelper) GetErrCount() int64 {
	return h.errCount
}

func (h *PasswdHelper) SetKey(key string, expiration time.Duration) {
	h.key = key
	h.expiration = expiration
}

func (h *PasswdHelper) fail(ctx context.Context) {
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

func (h *PasswdHelper) delIncr(ctx context.Context) {
	if err := h.rdb.Del(ctx, h.key).Err(); err != nil {
		h.OnErr(err)
	}
}