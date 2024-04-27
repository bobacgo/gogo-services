package security

import (
	"github.com/gogoclouds/gogo-services/framework/pkg/ucrypto"
)

// Ciphertext 密文
// use: 前端密码字段的传输
//
// 密码字段设计:
// 1.前端密码字段加密
// 2.后端解密出原文
// 3.后端密码强度校验
// 4.入库时hash不可逆编码(可以加盐)
type Ciphertext string

func (ct *Ciphertext) Decrypt(conf Config) error {
	cfg := conf.Ciphertext
	// TODO 密码强度规则校验
	if !cfg.IsCiphertext {
		return nil
	}
	pt, err := ucrypto.AESDecrypt(string(*ct), string(cfg.CipherKey))
	if err != nil {
		return err
	}
	*ct = Ciphertext(pt)
	return nil
}

func (ct *Ciphertext) BcryptHash() string {
	hash := DefaultPasswdVerifier.BcryptHash(string(*ct))
	return hash
}

func (ct *Ciphertext) BcryptVerify(hashPasswd string) bool {
	return DefaultPasswdVerifier.BcryptVerify(hashPasswd, string(*ct))
}
