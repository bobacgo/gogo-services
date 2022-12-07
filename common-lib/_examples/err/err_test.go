package err

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestErr(t *testing.T) {
	err := api()
	fmt.Printf("%+v", err)
}

func api() error {
	return biz()
}

func biz() error {
	return dao()
}

func dao() error {
	err := fmt.Errorf("bad parameter: %d", -1)
	return errors.WithStack(err)
	// or
	// return errors.New("data error")
}