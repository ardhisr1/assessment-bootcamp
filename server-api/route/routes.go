package route

import (
	"server-api/auth"
	"server-api/config"
	"server-api/handler"
	"server-api/user"

	"github.com/gin-gonic/gin"
)

var (
	DB              = config.Connection()
	userRepo        = user.NewRepository(DB)
	userService     = user.NewService(userRepo)
	userAuthService = auth.NewService()
	userHandler     = handler.NewUserHandler(userService, userAuthService)
)

func Routes(r *gin.Engine) {
	r.POST("/user/register", userHandler.RegisterUserHandler)
	r.POST("/user/login", userHandler.LoginUserHandler)
	r.GET("/user-detail", handler.Middleware(auth.NewService()), userHandler.GetUserDetailHandler)
}
