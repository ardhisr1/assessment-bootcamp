package main

import (
	"server-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.Routes(r)
	r.Run(":9000")
}
