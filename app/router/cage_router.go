package router

import (
	"JurrassicParkAPI/app/controller"
	"JurrassicParkAPI/app/pkg"
	log "github.com/sirupsen/logrus"
)

type CageRoutes struct {
	handler        *pkg.RequestHandler
	cageController *controller.CageController
}

// Setup user routes
func (c *CageRoutes) Setup() {
	log.Info("Setting up routes")

	api := c.handler.Gin.Group("/api")
	{
		api.GET("/cage", c.cageController.GetAllCages)
		api.GET("/cage/:id", c.cageController.GetOneCage)
		api.POST("/cage", c.cageController.CreateCage)
		api.PATCH("/cage/:id", c.cageController.UpdateCage)
	}
}

// NewCageRoutes  creates new crate controller
func NewCageRoutes(
	handler *pkg.RequestHandler,
	cageController *controller.CageController,
) *CageRoutes {
	return &CageRoutes{
		handler,
		cageController,
	}
}
