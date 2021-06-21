package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//routename(r)
	r.Run("${:9000}")
}
