package v1_test

import (
	"encoding/json"
	"fmt"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	//w := httptest.NewRecorder()
	request := v1.AdminLoginRequest{
		Username: "admin",
		Password: "",
	}
	reqData, _ := json.Marshal(request)
	reader := strings.NewReader(string(reqData))
	resp, err := http.Post("http://localhost:8080/admin/login", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	fmt.Printf("%s", response)
	//assert.Equal(t, 200, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}
