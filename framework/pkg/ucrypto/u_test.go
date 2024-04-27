package ucrypto_test

import (
	"encoding/base64"
	"testing"

	"github.com/gogoclouds/gogo-services/framework/pkg/ucrypto"
)

func TestBase64(t *testing.T) {
	raw := "hello world"
	t.Log(raw)
	ciphertext := base64.StdEncoding.EncodeToString([]byte(raw))
	t.Log(ciphertext)
	plaintext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", plaintext)
}

func TestAES(t *testing.T) {
	key := "1234567890123456" // 16、24、32 位
	raw := "hello world"
	t.Log(raw)
	ciphertext, err := ucrypto.AESEncrypt(raw, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ciphertext)
	plaintext, err := ucrypto.AESDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(plaintext)
}

func TestDES(t *testing.T) {
	key := "12345678" // 8 位
	raw := "hello world"
	t.Log(raw)
	ciphertext, err := ucrypto.DESEncrypt(raw, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ciphertext)
	plaintext, err := ucrypto.DESDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(plaintext)
}

func TestKeyPairs(t *testing.T) {
	ucrypto.KeyPairs("rsa")
}

// 私钥生成
// openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCuJb3nONLWxpn/jOo01wnDffefzdl/YYzZacR+tBgGFbAr4g1D
9nQjTUo1aahYCoMX1Ff8VqsE0tjA3+ZRL+dA7Y7E07+wBwGagXDqVNLEslnaBvKx
z/A0UCAKezVCQws9aSIdqujPm2pESKtJ+U6evogkCianKss6kbnbLnTfBwIDAQAB
AoGAIB432whA4nm4d0hO/bXXSCXYYLG/dl3Qc1ytb8zZTW38kutbFPjETKp5kEZP
VQWDTgbMv25glkAo19Gzka+rD6GeCqNkg1SkM9KSxM3/eS1CRIz2Sa9KeCljrG3e
xQI5MPKCUzFv0YjYmG1Egrm5uRjs5k1PXV1kQXNaJhhAJGECQQDUWlA2P4KZv9zF
9in/3EkHYojJzqur5cD8BB8WDi5vIlkV7tQQwayVR2sPOF1f4IQumFpaaB6navp0
gg08/Md3AkEA0fEenbr/4xGRiO6YwHcq76nznozJYooBuSPdo0+jqOkZnqMPPFju
W2HJ/8GruH3H0f8Yk6VQpWYUBrDxfuGo8QJBAIe2YGULGdBhChuKQzU99348vuca
qiRl5XwqtiNGVO65qO2XgPhkjoOo7QcBIsvPlSqiO7xjppOgjwg+xW8grekCQQCX
EdG9II33oHHAPijO/jF4SixTH+3eKX658dQQK0OSTUIxRBa3jyrduQ15K6zc0i3S
r6TIwcG5cy3f7r2oVsuRAkBhz2tIq5njpCMVm7tXVWXk5kH9p3g4TN/qqpEnOMQJ
xaV34mVJkwfzqDXpa6AvzeweCfRqndIcckgz8/jAYviu
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCuJb3nONLWxpn/jOo01wnDffef
zdl/YYzZacR+tBgGFbAr4g1D9nQjTUo1aahYCoMX1Ff8VqsE0tjA3+ZRL+dA7Y7E
07+wBwGagXDqVNLEslnaBvKxz/A0UCAKezVCQws9aSIdqujPm2pESKtJ+U6evogk
CianKss6kbnbLnTfBwIDAQAB
-----END RSA PUBLIC KEY-----
`)

func TestRSA(t *testing.T) {
	raw := "holle word"
	t.Log(raw)
	ciphertext, err := ucrypto.RSAEncrypt(publicKey, raw)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ciphertext)
	plaintext, err := ucrypto.RSADecrypt(privateKey, ciphertext)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(plaintext)
}

func TestMD5(t *testing.T) {
	data := "hello world"
	t.Log(data)
	md5 := ucrypto.MD5(data)
	t.Log(md5)
}

func TestSHA1(t *testing.T) {
	data := "hello world"
	t.Log(data)
	sha1 := ucrypto.SHA1(data)
	t.Log(sha1)
}

func TestSHA256(t *testing.T) {
	secret, data := "w3xeayw5smcn5ei0", "hello world"
	t.Logf("Secret: %s Data: %s\n", secret, data)
	sha := ucrypto.SHA256(secret, data)
	t.Log(sha)
}
