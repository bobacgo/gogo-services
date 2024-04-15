package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/server/http/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/web/gin/validator"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"

	"github.com/gogoclouds/gogo-services/common-lib/app/logger"

	"github.com/gin-gonic/gin"
)

func RunHttpServer(app *App, register func(e *gin.Engine, a *Options)) {
	app.wg.Add(1)
	defer app.wg.Done()

	e := gin.New()

	switch app.opts.conf.Env {
	case conf.EnvProd:
		gin.SetMode(gin.ReleaseMode)
	case conf.EnvDev:
		gin.SetMode(gin.DebugMode)
	case conf.EnvTest:
		gin.SetMode(gin.TestMode)
	}

	e.Use(gin.Logger()) // TODO -> zap.Logger
	e.Use(middleware.Recovery())
	e.Use(middleware.LoggerResponseFail())

	binding.Validator = new(validator.DefaultValidator)
	healthApi(e) // provide health API

	if register != nil {
		register(e, &app.opts) // register router
	}

	srv := &http.Server{Addr: app.opts.conf.Server.Http.Addr, Handler: e}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Panicf("listen: %s\n", err)
		}
	}()
	logger.Infof("http server running %s", app.opts.conf.Server.Http.Addr)
	<-app.exit
	logger.Info("Shutting down http server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Http server forced to shutdown: %w", err)
	}
	logger.Info("http server exiting")
}

// healthApi http check-up API
func healthApi(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		msg := fmt.Sprintf("%s [env=%s] %s, is active", conf.Conf.Name, conf.Conf.Env, conf.Conf.Version)
		r.Reply(c, msg)
	})
}
