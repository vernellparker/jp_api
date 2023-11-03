package repository

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewCageRepository),
	fx.Provide(NewDinosaurRepository),
	fx.Provide(NewVerificationRepository),
)
