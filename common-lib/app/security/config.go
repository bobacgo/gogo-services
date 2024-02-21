package security

import "time"

type Config struct {
	Secret              string `yaml:"secret"`              // jwt secret
	Issuer              string `yaml:"issuer"`              // jwt issuer
	AccessTokenExpired  string `yaml:"accessTokenExpired"`  // jwt access token expired
	RefreshTokenExpired string `yaml:"refreshTokenExpired"` // jwt refresh token expired
}

func (c *Config) GetAccessTokenExpired() time.Duration {
	d, _ := time.ParseDuration(c.AccessTokenExpired)
	return d
}

func (c *Config) GetRefreshTokenExpired() time.Duration {
	d, _ := time.ParseDuration(c.RefreshTokenExpired)
	return d
}

// TODO validate config

func (c *Config) Validate() []error {
	// TODO valid config data
	return nil
}