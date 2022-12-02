package server_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/gogoclouds/gogo-services/common-lib/internal/server"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"

	"github.com/gin-gonic/gin"
)

func Test_HttpServer(t *testing.T) {
	server.RunHttpServer(":8080", router)
}

func Test_HttpApi(t *testing.T) {
	r, err := http.DefaultClient.Get("http://127.0.0.1:8080/ping")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		t.Fatalf("http status: %s", r.Status)
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}

func router(h http.Handler) {
	e := h.(*gin.Engine)
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, r.Success())
	})
}
