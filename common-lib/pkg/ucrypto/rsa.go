package ucrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
	"os"
)

// RSA plaintext 不能太长

// 私钥生成
// openssl genrsa -out rsa_private_key.pem 1024

// 公钥: 根据私钥生成
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

// RSAEncrypt 加密
func RSAEncrypt(publicKey []byte, plaintext string) (string, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ct, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(plaintext))
	return base64.URLEncoding.EncodeToString(ct), err
}

// RSADecrypt 解密
func RSADecrypt(privateKey []byte, ciphertext string) (string, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	cth, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	pt, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cth)
	return string(pt), err
}

// KeyPairs 生成RSA密钥对
func KeyPairs(keyName string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}
	x509Encoded := x509.MarshalPKCS1PrivateKey(privateKey)
	savePEMKey(keyName+".private.pem", &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(privateKey.Public())
	savePEMKey(keyName+".public.pem", &pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509EncodedPub})
}

// savePEMKey 保存PEM格式的密钥到文件
func savePEMKey(fileName string, block *pem.Block) {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = pem.Encode(outFile, block)
	if err != nil {
		log.Fatal(err)
	}
}
