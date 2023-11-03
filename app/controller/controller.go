package controller

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewCageController),
	fx.Provide(NewDinosaurController),
)
