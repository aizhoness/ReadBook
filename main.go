package main

import (
	"ReadBook/logger"
	handler "ReadBook/route_handler"
	"ReadBook/router"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			logger.Options,
			router.Options,
		),
		fx.Invoke(Register),
	).Run()
}

func Register(
	lc fx.Lifecycle,
	logger *zap.SugaredLogger,
	router *mux.Router,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				addr := ":8080"
				logger.Info("Listening on ", addr)
				go http.ListenAndServe(addr, router)
				return nil
			},
			OnStop: func(context.Context) error {
				defer logger.Sync()
				return nil
			},
		},
	)

	handler.New(logger, router)

}
