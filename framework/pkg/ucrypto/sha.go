package ucrypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// --- SHA签名 ---

func SHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	// fmt.Sprintf("%x", h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func SHA256(secret, data string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write([]byte(data))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
