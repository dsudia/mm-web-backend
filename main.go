package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMainEngine() *gin.Engine {
	r := gin.Default()

	// Healthchecker for ELB
	r.GET("/check", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
		c.AbortWithStatus(200)
	})

	r.POST("/signup", func(c *gin.Context) {
		var ne NewEducator
		c.BindJSON(&ne)

		h := bcrypt.GenerateFromPassword(ne.Password, 10)

		ne.Password = h
	})
}

func main() {
	getMainEngine().Run()
}
