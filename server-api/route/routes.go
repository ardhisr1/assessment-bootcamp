package route

import (
	"server-api/config"
	"server-api/handler"
	"server-api/user"

	"github.com/gin-gonic/gin"
)

var (
	DB          = config.Connection()
	UserRepo    = user.NewRepository(DB)
	UserService = user.NewService(UserRepo)
	userHandler = handler.NewHandler(UserService)
)

func route(r *gin.Engine) {
	r.POST("/user/register", handler.RegisterUser)
}
