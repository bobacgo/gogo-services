package uid_test

import (
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uid"
	"testing"
)

func TestRandSeqID(t *testing.T) {
	randString := uid.RandSeqID(16)
	for i := 0; i < 10; i++ {
		t.Log(randString())
	}
}
