package services

import (
  "fmt"
  "strings"
 // "github.com/gin-gonic/gin"
)

type CharacterInputs struct {
  IsTalking bool   `json: "isTalking"`
  Action    string `json: "action"`
  Emotion   string `json: "emotion"`
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

// Create a prompt string for a character in a panel based on user inputs
func buildCharacterPrompt(name string, character CharacterInputs) string {
  Talking := map[string]string {
    "Talking": "is talking with mouth open",
    "NotTalking": "is not talking, mouth closed",
  }

  talking := Talking["NotTalking"]
  if character.IsTalking {
    talking= Talking["Talking"]
  }

  action := character.Action
  emotion := character.Emotion

  prompt := fmt.Sprintf(
    "%s %s. %s is feeling %s. %s %s",
    name, action, name, emotion, name, talking,
  )

  return prompt
}

// Create a prompt string for a character's physical description
// in a panel based on predefined traits
func buildCharacterDescription(name string) string {
  // TODO mutex lock these maps as globals?
  // TODO manage character definitions elsewhere
  Character := map[string]CharacterTraits {
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

  character := Character[name]

  prompt := fmt.Sprintf(
    "%s %s. %s is wearing %s.",
    name, character.Appearance, name, character.Outfit,
  )

  return prompt
}

// Create an image gen text prompt
func AssembleImagePrompt(promptInput PromptInput) (string, error) {
  // TODO mutex lock these maps as globals?
  Color := map[string]string {
    "Color": "full color",
    "Grayscale": "black-and-white",
  }

  Style := map[string]string {
    "Josei": "josei",
    "Seinen": "seinen",
    "SliceOfLife": "slice-of-life",
    "Shonen": "shonen manga style",
    "Shojo": "shojo manga style",
  }

  CameraAngle := map[string]string {
    "BirdsEye": "bird's-eye view",
    "CloseUp": "close-up shot",
    "TwoShot": "two-shot",
    "EyeLevel": "eye level shot",
    "Establishing": "establishing shot",
  }

  var characterPrompts, characterDescriptions []string

  for name, character := range promptInput.Characters {
    characterPrompt := buildCharacterPrompt(name, character)
    characterPrompts = append(characterPrompts, characterPrompt)

    characterDesc := buildCharacterDescription(name)
    characterDescriptions = append(characterDescriptions, characterDesc)
  }

  joinedPrompts := strings.Join(characterPrompts, "\n")
  joinedDescs := strings.Join(characterDescriptions, "\n")

  imagePrompt := fmt.Sprintf(
    "Draw a single manga panel from a %s camera angle in a %s, %s manga style. It should be a single panel with no panel dividers.\n\n%s\n\nThe setting is: %s\n%s",
    CameraAngle[promptInput.CameraAngle], Color[promptInput.Color],
    Style[promptInput.Style], joinedDescs, promptInput.Environment,
    joinedPrompts,
  )
  fmt.Println(imagePrompt)

  return imagePrompt, nil
}

func GenerateImage(prompt string) (string, error) {
  fmt.Println("BBBBB")
  return "B", nil
}
