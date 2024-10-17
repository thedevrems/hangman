package main

import (
	"hangman/configuration"
	"hangman/managegame"
)

func main() {
	// Configuration Accounting
	dataConfig := configuration.DefaultConfiguration()
	dataFiles := configuration.ConfigFiles()
	dataGame := configuration.ConfigGame()
	dataError := configuration.ConfigError()
	// Selection of language
	var dataTranslation *configuration.Structure_translation
	if dataConfig.Language == dataGame.English {
		dataTranslation = configuration.DataTranslationEn()
	} else {
		dataTranslation = configuration.DataTranslationFr()
	}

	// Start Game
	managegame.Panel(dataConfig, dataGame, dataError, dataTranslation, dataFiles)
}
