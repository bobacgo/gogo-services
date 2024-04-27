package logger

import (
	"log/slog"

	slogcontext "github.com/PumpkinSeed/slog-context"
	slogmulti "github.com/samber/slog-multi"
)

func InitSlog(handlers ...slog.Handler) {
	sc := slogcontext.NewHandler(slogmulti.Fanout(handlers...))
	l := slog.New(sc)
	slog.SetDefault(l)
}
