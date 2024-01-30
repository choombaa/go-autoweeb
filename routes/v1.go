package routes

import (
  "github.com/gin-gonic/gin"
  "go-autoweeb/controllers"
)

func LoadV1Routes(router *gin.Engine) {
  // Generate an image
  router.POST("/image", func(ctx *gin.Context) {
    controllers.GenerateImage(ctx)
  })
}
