package ucrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// --- AES加密解密(对称) ---

/*
	前端代码实现

	import CryptoJS from 'crypto-js'

	aseEncrypt()

	function aseEncrypt() {
		let msg = "hello world"
		let key = CryptoJS.enc.Utf8.parse("1234567890123456")
		var encrypted = CryptoJS.AES.encrypt(msg, key, {
			iv: key,
			mode: CryptoJS.mode.CBC,// CBC算法
			padding: CryptoJS.pad.Pkcs7 //使用pkcs7 进行padding 后端需要注意
		})
		let text = encrypted.ciphertext.toString(CryptoJS.enc.Hex)
		console.log(text)
	}

*/

func AESEncrypt(plaintext, key string) (string, error) {
	pb, kb := []byte(plaintext), []byte(key)
	// 秘钥块
	bl, err := aes.NewCipher(kb)
	if err != nil {
		return "", fmt.Errorf("填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256  key 长度必须 16|24|32长度: %s", err)
	}
	// 补全码
	pb = pkcs7Padding(pb, bl.BlockSize())
	// 加密模式
	bm := cipher.NewCBCEncrypter(bl, kb)
	// 创建数组
	cpb := make([]byte, len(pb))
	// 加密
	bm.CryptBlocks(cpb, pb)
	return hex.EncodeToString(cpb), nil
}

func AESDecrypt(ciphertext, key string) (pt string, err error) {
	// panic recover bm.CryptBlocks
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	ctb, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	kb := []byte(key)
	bl, err := aes.NewCipher(kb) //选择加密算法
	if err != nil {
		return "", fmt.Errorf("key 长度必须 16|24|32长度: %s", err)
	}
	bm := cipher.NewCBCDecrypter(bl, kb) // 前端代码对应: mode: CryptoJS.mode.CBC,// CBC算法
	ptb := make([]byte, len(ctb))
	bm.CryptBlocks(ptb, ctb)
	ptb = pkcs7UnPadding(ptb) // 前端代码对应:  padding: CryptoJS.pad.Pkcs7
	return string(ptb), nil
}

// PKCS7Padding 补码
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	p := blockSize - len(ciphertext)%blockSize
	ptb := bytes.Repeat([]byte{byte(p)}, p)
	return append(ciphertext, ptb...)
}

// pkcs7UnPadding 去补码
func pkcs7UnPadding(plaintext []byte) []byte {
	l := len(plaintext)
	unp := int(plaintext[l-1])
	return plaintext[:(l - unp)]
}
