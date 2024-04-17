package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/server/http/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"

	"github.com/gogoclouds/gogo-services/common-lib/app/logger"

	"github.com/gin-gonic/gin"
)

func RunHttpServer(app *App, register func(e *gin.Engine, a *Options)) {
	app.wg.Add(1)
	defer app.wg.Done()

	cfg := app.opts.GetConf()

	switch cfg.Env {
	case conf.EnvProd:
		gin.SetMode(gin.ReleaseMode)
	case conf.EnvDev:
		gin.SetMode(gin.DebugMode)
	case conf.EnvTest:
		gin.SetMode(gin.TestMode)
	}

	e := gin.New()

	e.Use(gin.Logger()) // TODO -> zap.Logger
	e.Use(middleware.Recovery())
	e.Use(middleware.LoggerResponseFail())

	if cfg.Env != conf.EnvDev {
		slog.Warn(fmt.Sprintf(`[gin] Running in "%s" mode`, gin.Mode()))
	}

	healthApi(e, cfg) // provide health API

	if register != nil {
		register(e, &app.opts) // register router
	}

	srv := &http.Server{Addr: cfg.Server.Http.Addr, Handler: e}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Panicf("listen: %s\n", err)
		}
	}()
	logger.Infof("http server running %s", cfg.Server.Http.Addr)
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
func healthApi(e *gin.Engine, cfg *conf.BasicConfig) {
	e.GET("/health", func(c *gin.Context) {
		msg := fmt.Sprintf("%s [env=%s] %s, is active", cfg.Name, cfg.Env, cfg.Version)
		r.Reply(c, msg)
	})
}
