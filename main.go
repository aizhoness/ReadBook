package main

import (
	"ReadBook/config"
	"ReadBook/dbpart"
	"ReadBook/logger"
	handler "ReadBook/route_handler"
	"ReadBook/router"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			logger.Options,
			router.Options,
			config.Options,
			dbpart.Options,
		),
		fx.Invoke(Register),
	).Run()
}

func Register(
	lc fx.Lifecycle,
	logger *zap.SugaredLogger,
	router *mux.Router,
	config *config.Config,
	database *sqlx.DB,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				addr := config.Server.Host + ":" + config.Server.Port
				logger.Info("Listening on ", addr)
				go http.ListenAndServe(addr, router)
				return nil
			},
			OnStop: func(context.Context) error {
				defer logger.Sync()
				defer database.Close()
				return nil
			},
		},
	)

	handler.New(logger, router, database)

}
