package route

import (
	"server-api/auth"
	"server-api/config"
	"server-api/handler"
	"server-api/user"
)

var (
	DB              = config.Connection()
	userRepo        = user.NewRepository(DB)
	userService     = user.NewService(userRepo)
	userAuthService = auth.NewService()
	userHandler     = handler.NewUserHandler(userService, userAuthService)
)
