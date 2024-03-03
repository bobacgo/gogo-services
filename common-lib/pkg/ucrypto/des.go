package ucrypto

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

// --- DES加密解密(对称) ---

func DESEncrypt(text, key string) (string, error) {
	src, kb := []byte(text), []byte(key)
	block, err := des.NewCipher(kb)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = zeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("need a multiple of the block size")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DESDecrypt(decrypted, key string) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	kb := []byte(key)
	block, err := des.NewCipher(kb)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = zeroUnPadding(out)
	return string(out), nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
