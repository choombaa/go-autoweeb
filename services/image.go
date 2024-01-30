package services

import (
  "fmt"
 // "github.com/gin-gonic/gin"
)

type CharacterInputs struct {
  IsTalking bool `json: "isTalking"`
}

type PromptInput struct {
  CameraAngle string   `json: "cameraAngle" binding: "required"`
  Color       string   `json: "color"       binding: "required"`
  Style       string   `json: "style"       binding: "required"`
  Environment string   `json: "environment  binding: "required"`
  Characters  map[string]CharacterInputs `json: "characters"`
}

type CharacterTraits struct {
  Appearance string `json: "appearance" binding: "required"`
  Outfit     string `json: "outfit"     binding: "required"`
}

func AssembleImagePrompt(promptInput PromptInput) (string, error) {
  // TODO mutex lock these maps as globals?
  //Color := map[string]string {
  _ = map[string]string {
    "Color": "full color",
    "Grayscale": "black-and-white",
  }

  //Style := map[string]string {
  _ = map[string]string {
    "Josei": "josei",
    "Seinen": "seinen",
    "SliceOfLife": "slice-of-life",
    "Shonen": "shonen manga style",
    "Shojo": "shojo manga style",
  }

  //CameraAngle := map[string]string {
  _ = map[string]string {
    "BirdsEye": "bird's-eye view",
    "CloseUp": "close-up shot",
    "TwoShot": "two-shot",
    "EyeLevel": "eye level shot",
    "Establishing": "establishing shot",
  }

  //Talking := map[string]string {
  _ = map[string]string {
    "Talking": "is talking with mouth open",
    "NotTalking": "is not talking, mouth closed",
  }

  //Character := map[string]CharacterTraits {
  _ = map[string]CharacterTraits {
    "Danielle": {
      Appearance: "is a beautiful female tennis player. She is black. She has a medium build and a brown ponytail.",
      Outfit: "white tennis skirt, blue polo shirt, light blue baseball cap",
    },
    "Haruto": {
      Appearance: "looks like a typical Japanese manga protagonist. He is Japanese. He has a medium build and long black hair that falls down on his face.",
      Outfit: "light gray hoodie",
    },
    "Tsunemori": {
      Appearance: "is a tall, beautiful Japanese anime woman. She is curvy and has light brown twintails.",
      Outfit: "crop-top-style hoodie, with camouflage cargo pants. She wears an eyepatch over her right eye.",
    },
  }

  var includedCharacters []string
  for name := range promptInput.Characters {
    fmt.Println(name)
    includedCharacters = append(includedCharacters, name)
  }

  fmt.Println(includedCharacters)
  return "A", nil
}

func GenerateImage() (string, error) {
  fmt.Println("BBBBB")
  return "B", nil
}
