package app_test

import (
	"io"
	"net/http"
	"testing"

	app "github.com/gogoclouds/gogo-services/framework/app"

	"github.com/gin-gonic/gin"
)

func Test_HttpServer(t *testing.T) {
	app.RunMustHttpServer(nil, router)
}

func Test_HttpApi(t *testing.T) {
	r, err := http.DefaultClient.Get("http://127.0.0.1:8080/ping")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		t.Fatalf("http status: %s", r.Status)
	}
	b, err2 := io.ReadAll(r.Body)
	if err2 != nil {
		t.Fatal(err2)
	}
	t.Logf("%s", b)
}

func router(e *gin.Engine, app *app.Options) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		})
	})
}
