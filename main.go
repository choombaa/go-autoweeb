package main

import (
  "net/http"
  "go-autoweeb/routes"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

func setupRouter() *gin.Engine {
  // Disable Console Color
  // gin.DisableConsoleColor()
  router := gin.Default()
  // TODO Default CORS doesn't play nice with Authorization header
  router.Use(cors.Default())

  routes.LoadV1Routes(router);

  return router
}

func main() {
  router := setupRouter()
  // Listen and Server in 0.0.0.0:8080
  router.Run(":8080")
}
