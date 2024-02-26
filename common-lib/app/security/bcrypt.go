package security

import (
	"github.com/gogoclouds/gogo-services/common-lib/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

var bcryptSalt = util.RandSeqID(8)

// BcryptHash 明文加密
func BcryptHash(passwd string) (hash, salt string) {
	salt = bcryptSalt()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(passwd+salt), bcrypt.DefaultCost)
	return string(bytes), salt
}

// BcryptVerify 校验密文和明文
func BcryptVerify(salt, hash, passwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd+salt))
	return err == nil
}