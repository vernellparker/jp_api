package bootstrap

import (
	"JurrassicParkAPI/app/pkg"
	"JurrassicParkAPI/app/router"
	"JurrassicParkAPI/config"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
)

// registerHooks is used to start and stop the service as it is invoked by Uber Fx
func registerHooks(lifecycle fx.Lifecycle, h *pkg.RequestHandler, routes *router.Routes, db *config.Database) {
	routes.Setup()
	portNumber := os.Getenv("PORT")
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					err := h.Gin.Run(fmt.Sprintf(":%s", portNumber))
					if err != nil {
						log.Fatal(err)
					}
				}()
				return nil
			},

			OnStop: func(context.Context) error {
				log.Info("Stopping application")
				err := db.DB.Close()
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	)
}
