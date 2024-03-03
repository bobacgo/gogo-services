package config

import "time"

type Config struct {
	Ciphertext CiphertextConfig `mapstructure:"ciphertext"`
	Jwt        JwtConfig        `mapstructure:"jwt"`
}

type CiphertextConfig struct {
	IsCiphertext bool   `mapstructure:"isCiphertext"` // 密码字段是否启用密文传输
	CipherKey    string `mapstructure:"cipherKey"`    // 支持 8 16 24 bit
}

type JwtConfig struct {
	Secret              string `mapstructure:"secret"`              // jwt secret
	Issuer              string `mapstructure:"issuer"`              // jwt issuer
	AccessTokenExpired  string `mapstructure:"accessTokenExpired"`  // jwt access token expired
	RefreshTokenExpired string `mapstructure:"refreshTokenExpired"` // jwt refresh token expired
}

func (c *JwtConfig) GetAccessTokenExpired() time.Duration {
	d, _ := time.ParseDuration(c.AccessTokenExpired)
	return d
}

func (c *JwtConfig) GetRefreshTokenExpired() time.Duration {
	d, _ := time.ParseDuration(c.RefreshTokenExpired)
	return d
}

// TODO validate config

func (c *JwtConfig) Validate() []error {
	// TODO valid config data
	return nil
}
