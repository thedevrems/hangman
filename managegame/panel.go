package managegame

import (
	"fmt"
	"strings"

	"github.com/thedevrems/hangman/manageerror"

	"github.com/thedevrems/hangman/managefiles"

	"github.com/thedevrems/hangman/configuration"
)

// Panel manages the game's configuration panel, allowing the user to modify
// certain game settings before starting a new game.
//
// Parameters:
//   - dataConfig (*configuration.Structure_configuration): structure containing the current game configuration settings (language, victory counter, difficulty, etc.).
//   - dataGame (*configuration.Structure_game): structure containing game-specific information (available languages, etc.).
//   - dataError (*configuration.Structure_error): structure containing error handling information (error messages, etc.).
//   - dataTranslation (*configuration.Structure_translation): structure containing translations of the messages displayed according to the selected language.
//   - dataFiles (*configuration.Structure_files): structure containing information about the files used in the game.
//
// Features:
//
//	The function runs an infinite loop, allowing the user to modify game settings such as:
//	  - Game language (option 1).
//	  - Enabling or disabling the victory counter (option 2).
//	  - Enabling or disabling difficulty management (option 3).
//	  - Allowing or disallowing the addition of a word after each game (option 4).
//	The user can choose to start the game (option 5) after applying the desired configuration.
//
// Main steps:
//  1. Displays the current configuration panel, showing the state of each option (language, victory counter, etc.).
//  2. Allows the user to select an option by entering a number.
//  3. Based on the selected number, the panel allows modification of a particular option or starting the game.
//  4. Handles errors in case of incorrect input by displaying an error message via `manageerror.PrintError`.
//
// Available options:
//  1. Choose the game language (English or French).
//  2. Enable or disable the victory counter.
//  3. Enable or disable difficulty management.
//  4. Allow or disallow adding a new word after each game.
//  5. Start the game with the current configuration.
//  6. Leave the game
func Panel(dataConfig *configuration.Structure_configuration, dataGame *configuration.Structure_game, dataError *configuration.Structure_error, dataTranslation *configuration.TranslationHangman, dataFiles *configuration.Structure_files) {
	for {
		ClearScreen()
		var stringVictoryCounter string
		var stringEnableDifficulty string
		var stringAddWordAfterGame string
		var selectMod string
		var stringEnableJokers string

		// Determine the current state of the victory counter, difficulty, and add word after game
		if dataConfig.VictoryCounter {
			stringVictoryCounter = dataTranslation.Enable
		} else {
			stringVictoryCounter = dataTranslation.Disable
		}

		if dataConfig.EnableDifficulty {
			stringEnableDifficulty = dataTranslation.Enable
		} else {
			stringEnableDifficulty = dataTranslation.Disable
		}

		if dataConfig.AddWordAfterGame {
			stringAddWordAfterGame = dataTranslation.Enable
		} else {
			stringAddWordAfterGame = dataTranslation.Disable
		}

		if dataConfig.EnableJokers {
			stringEnableJokers = dataTranslation.Enable
		} else {
			stringEnableJokers = dataTranslation.Disable
		}

		// Display the current configuration and available options
		fmt.Println(dataTranslation.ConfigurationCurrent)
		fmt.Println(dataTranslation.EnterNumber)
		fmt.Println(dataTranslation.OptionLanguage + " (" + dataGame.English + "/" + dataGame.Frensh + ") : " + dataConfig.Language)
		fmt.Println(dataTranslation.OptionVictoryCounter + stringVictoryCounter)
		fmt.Println(dataTranslation.OptionEnableDifficulty + stringEnableDifficulty)
		fmt.Println(dataTranslation.OptionAddWordAfterGame + stringAddWordAfterGame)
		fmt.Println(dataTranslation.OptionEnableJokers + stringEnableJokers)
		fmt.Println(dataTranslation.OptionStartGame)
		fmt.Println(dataTranslation.LeaveTheGame)
		fmt.Scanln(&selectMod)

		// Handle user-selected options
		if selectMod == "1" {
			ClearScreen()
			// Change game language
			for {
				var selectLang string
				fmt.Println(dataTranslation.PleaseEnterLanguage + " (" + dataGame.English + "/" + dataGame.Frensh + ") : ")
				fmt.Scanln(&selectLang)
				selectLang = strings.ToLower(selectLang)

				if selectLang == dataGame.English {
					dataConfig.Language = dataGame.English
					dataTranslation = configuration.DataTranslationHangmanEn()
					break
				} else if selectLang == dataGame.Frensh {
					dataConfig.Language = dataGame.Frensh
					dataTranslation = configuration.DataTranslationHangmanFr()
					break
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}

		} else if selectMod == "2" {
			ClearScreen()
			// Enable/disable victory counter
			for {
				var selectVictoryCounter string
				fmt.Println(dataTranslation.EnableWinCounter + "(" + dataTranslation.SelectModYes + "/" + dataTranslation.SelectModNo + ") :")
				fmt.Scanln(&selectVictoryCounter)
				selectVictoryCounter = strings.ToLower(selectVictoryCounter)

				if selectVictoryCounter == dataTranslation.SelectModYes {
					dataConfig.VictoryCounter = true
					break
				} else if selectVictoryCounter == dataTranslation.SelectModNo {
					dataConfig.VictoryCounter = false
					break
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}

		} else if selectMod == "3" {
			ClearScreen()
			// Enable/disable difficulty setting
			for {
				var selectDifficulty string
				fmt.Println(dataTranslation.EnableDifficulty + "(" + dataTranslation.SelectModYes + "/" + dataTranslation.SelectModNo + ") :")
				fmt.Scanln(&selectDifficulty)
				selectDifficulty = strings.ToLower(selectDifficulty)

				if selectDifficulty == dataTranslation.SelectModYes {
					dataConfig.EnableDifficulty = true
					break
				} else if selectDifficulty == dataTranslation.SelectModNo {
					dataConfig.EnableDifficulty = false
					break
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}

		} else if selectMod == "4" {
			ClearScreen()
			// Enable/disable adding a word after the game
			for {
				var selectAddWord string
				fmt.Println(dataTranslation.AllowAddingWordAtEnd + "(" + dataTranslation.SelectModYes + "/" + dataTranslation.SelectModNo + ") :")
				fmt.Scanln(&selectAddWord)
				selectAddWord = strings.ToLower(selectAddWord)

				if selectAddWord == dataTranslation.SelectModYes {
					dataConfig.AddWordAfterGame = true
					break
				} else if selectAddWord == dataTranslation.SelectModNo {
					dataConfig.AddWordAfterGame = false
					break
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}

		} else if selectMod == "5" {
			ClearScreen()
			// Enable/disable adding a word after the game
			for {
				var selectEnableJokers string
				fmt.Println(dataTranslation.EnableJokers + "(" + dataTranslation.SelectModYes + "/" + dataTranslation.SelectModNo + ") :")
				fmt.Scanln(&selectEnableJokers)
				selectEnableJokers = strings.ToLower(selectEnableJokers)

				if selectEnableJokers == dataTranslation.SelectModYes {
					dataConfig.EnableJokers = true
					break
				} else if selectEnableJokers == dataTranslation.SelectModNo {
					dataConfig.EnableJokers = false
					break
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}

		} else if selectMod == "6" {
			ClearScreen()
			// Creation of the environment
			managefiles.CaseManagementFiles(dataFiles, dataConfig, dataGame, dataError, dataTranslation)
			// Start the game after configuration
			StartGame(dataConfig, dataGame, dataError, dataTranslation, dataFiles)
		} else if selectMod == "7" {
			// Leave the game
			ClearScreen()
			break
		} else {
			ClearScreen()
			manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
		}
	}
}
