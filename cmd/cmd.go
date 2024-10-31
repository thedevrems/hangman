package main

import (
	"github.com/thedevrems/hangman/configuration"

	"github.com/thedevrems/hangman/managegame"
)

func main() {
	// Configuration Accounting
	dataConfig := configuration.DefaultConfiguration()
	dataFiles := configuration.ConfigFiles()
	dataGame := configuration.ConfigGame()
	dataError := configuration.ConfigError()
	// Selection of language
	var dataTranslation *configuration.TranslationHangman
	if dataConfig.Language == dataGame.English {
		dataTranslation = configuration.DataTranslationHangmanEn()
	} else {
		dataTranslation = configuration.DataTranslationHangmanFr()
	}

	// Start Game
	managegame.Panel(dataConfig, dataGame, dataError, dataTranslation, dataFiles)
}
