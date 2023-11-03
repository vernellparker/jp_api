package router

import (
	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewCageRoutes),
	fx.Provide(NewDinosaurRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	cageRoutes *CageRoutes,
	dinosaurRoutes *DinosaurRoutes,
) *Routes {
	return &Routes{
		cageRoutes,
		dinosaurRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
