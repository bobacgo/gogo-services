package security

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"time"
)

const (
	ATokenExpiredDuration = 2 * time.Hour
	RTokenExpiredDuration = 30 * 24 * time.Hour
)

type JWToken struct {
	SigningKey          []byte
	Issuer              string        `yaml:"issuer"`              // jwt issuer
	AccessTokenExpired  time.Duration `yaml:"accessTokenExpired"`  // jwt access token expired
	RefreshTokenExpired time.Duration `yaml:"refreshTokenExpired"` // jwt refresh token expired
}

func NewJWT(conf *Config) *JWToken {
	return &JWToken{
		SigningKey:          []byte(conf.Secret),
		Issuer:              conf.Issuer,
		AccessTokenExpired:  conf.GetAccessTokenExpired(),
		RefreshTokenExpired: conf.GetRefreshTokenExpired(),
	}
}

// Generate 颁发token access token 和 refresh token
func (t *JWToken) Generate(claims Claims) (atoken, rtoken string, err error) {
	if claims.ExpiresAt == 0 {
		if t.AccessTokenExpired == 0 {
			t.AccessTokenExpired = ATokenExpiredDuration
		}
		claims.ExpiresAt = time.Now().Add(t.AccessTokenExpired).Unix()
	}
	if t.Issuer != "" {
		claims.Issuer = t.Issuer
	}

	claims.NotBefore = time.Now().Unix() // 签名生效时间
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(t.SigningKey)
	if err != nil {
		return
	}
	// refresh token 不需要保存任何用户信息
	sc := claims.StandardClaims
	if t.RefreshTokenExpired == 0 {
		t.RefreshTokenExpired = RTokenExpiredDuration
	}
	sc.ExpiresAt = time.Now().Add(t.RefreshTokenExpired).Unix()
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, sc).SignedString(t.SigningKey)
	return
}

func (t *JWToken) keyfunc(_ *jwt.Token) (any, error) {
	return t.SigningKey, nil
}

// Verify 验证Token
func (t *JWToken) Verify(tokenString string) (*Claims, error) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, t.keyfunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token 验证不通过")
	}
	return claims, nil
}

// Refresh 通过 refresh token 刷新 atoken
func (t *JWToken) Refresh(atoken, rtoken string) (newAToken, newRToken, key string, err error) {
	// rtoken 无效直接返回
	if _, err = jwt.Parse(rtoken, t.keyfunc); err != nil {
		return
	}
	// 从旧access token 中解析出claims数据
	claim := new(Claims)
	_, err = jwt.ParseWithClaims(atoken, claim, t.keyfunc)
	// 判断错误是不是因为access token 正常过期导致的
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		at, rt, err := t.Generate(*claim)
		return at, rt, key, err
	}
	return
}