package security

import "testing"

func TestPwd(t *testing.T) {
	pv := new(PasswdVerifier)
	hash := pv.BcryptHash("admin123")
	t.Log(hash)
	if pv.BcryptVerify(hash, "admin123") {
		t.Log("success")
	} else {
		t.Log("fail")
	}
}