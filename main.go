package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMainEngine() *gin.Engine {
	r := gin.Default()

	// Healthchecker for ELB
	r.GET("/check", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
		c.AbortWithStatus(200)
	})

	r.POST("/signup/educator", routes.CreateEducatorRoute(*gin.Context))

	r.POST("/signup/school", routes.CreateSchoolRoute(*gin.Context))
}

func main() {
	getMainEngine().Run()
}
