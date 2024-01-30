package controllers

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "go-autoweeb/services"
)

func GenerateImage(ctx *gin.Context) {
  // Validate the request
  var promptInput services.PromptInput
  var validationErr = ctx.ShouldBindJSON(&promptInput)
  if validationErr != nil {
    fmt.Println("Error:", validationErr.Error())
    ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
    return
  }

  doneChannel := make(chan bool)
  var imagePrompt, image string
  var promptErr, genErr error

  // Assemble prompt
  go func() {
    imagePrompt, promptErr = services.AssembleImagePrompt(promptInput)
    doneChannel <- true
  }()
  <-doneChannel

  if promptErr != nil {
    fmt.Println("Error:", promptErr)
    ctx.JSON(http.StatusBadRequest, gin.H{"error": promptErr.Error()})
    return
  }
  fmt.Println(imagePrompt)

  // Generate image
  go func() {
    image, genErr = services.GenerateImage()
    doneChannel <- true
  }()
  <-doneChannel

  if genErr != nil {
    fmt.Println("Error:", genErr)
    ctx.JSON(http.StatusBadRequest, gin.H{"error": genErr.Error()})
    return
  }
  fmt.Println(image)

  ctx.JSON(http.StatusOK, gin.H{
    "message": "test",
  })
}
