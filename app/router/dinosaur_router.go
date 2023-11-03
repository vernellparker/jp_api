package router

import (
	"JurrassicParkAPI/app/controller"
	"JurrassicParkAPI/app/pkg"
	log "github.com/sirupsen/logrus"
)

type DinosaurRoutes struct {
	handler            *pkg.RequestHandler
	dinosaurController *controller.DinosaurController
}

// Setup user routes
func (c *DinosaurRoutes) Setup() {
	log.Info("Setting up routes")

	api := c.handler.Gin.Group("/api")
	{
		api.GET("/dino/preload", c.dinosaurController.PreloadDinosaurs)
		api.GET("/dino", c.dinosaurController.GetAllDinosaurs)
		api.GET("/dino/:id", c.dinosaurController.GetOneDinosaur)
		api.POST("/dino", c.dinosaurController.CreateDinosaur)
		api.PATCH("/dino/:id", c.dinosaurController.UpdateDinosaur)
	}
}

// NewDinosaurRoutes  creates new crate controller
func NewDinosaurRoutes(
	handler *pkg.RequestHandler,
	dinosaurController *controller.DinosaurController,
) *DinosaurRoutes {
	return &DinosaurRoutes{
		handler,
		dinosaurController,
	}
}
