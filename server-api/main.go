package main

import (
	"server-api/handler"
	"server-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(handler.CORSMiddleware())
	route.Routes(r)
	r.Run(":9000")
}
