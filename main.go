package main

import (
	"net/http"
  "go-autoweeb/routes"
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
  // TODO Default CORS doesn't play nice with Authorization header
  router.Use(cors.Default())

  routes.LoadV1Routes(router);

	// Get user value
	router.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
