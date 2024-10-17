package managegame

import (
	"fmt"
	"hangman/configuration"
	"hangman/manageerror"
	"strconv"
	"strings"
)

// PlayHangman starts the hangman game with support for other functions
// The explanation part by part is made throughout this long function
//
// Args :
// Takes a pointer to a Structure_error object as an argument.
// Takes a pointer to a Structure_translation object in the second argument
// Takes a pointer to a Structure_files object in the third argument
// Takes a pointer to a Structure_configuration object in the fourth argument
// Takes a pointer to a Structure_game object in the fifth argument
// Takes as argument the win number with an int type (if option was chosen)
// Takes as argument the number of defeats with an int type (if option was chosen)
// Also takes the difficulty level (if option was chosen)
//
// Return :
// The generic function first returns a win or lose message with a string type
// and also returns whether it is a win or not in boolean form
func PlayHangman(dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, dataFiles *configuration.Structure_files, dataConfig *configuration.Structure_configuration, dataGame *configuration.Structure_game, nbrVictory int, nbrLoose int, nameDifficulty string) (message string, isVictory bool) {
	var actuelWord string = ""
	var stageHangman int = 1
	var tryAgain string = ""
	var selectedLetter string = ""
	var tabSelectedLetter []string
	var tabSelectedWord []string
	var printSelected string
	var youWin bool = false
	var revealWord bool = false
	var wordsToGuess string = ChooseWord(dataError, dataTranslation, dataFiles, dataConfig, dataGame, nameDifficulty)

	if wordsToGuess == dataError.ErrSelectMode || wordsToGuess == dataError.ErrSelectWord {
		return wordsToGuess, false
	}

	// Indicates the number of Winnings and the number of Losses (if the extension was chosen by the player) but only if at least one game has been completed.
	fmt.Println(dataTranslation.TitleHangman)
	if nbrLoose != 0 || nbrVictory != 0 {
		fmt.Println(dataTranslation.YourCurrentlyOn + strconv.Itoa(nbrVictory) + dataTranslation.Victory + dataTranslation.And + strconv.Itoa(nbrLoose) + dataTranslation.Defeat)
	}

	for {
		// Show already selected letters
		if len(tabSelectedLetter) != 0 {
			printSelected = ""
			fmt.Println(dataTranslation.LettersAlreadyUsed)
			for i := 0; i < len(tabSelectedLetter); i++ {
				if i != len(tabSelectedLetter)-1 {
					printSelected += tabSelectedLetter[i] + " "
				} else {
					printSelected += tabSelectedLetter[i]
				}
			}
			fmt.Println(printSelected)
		}

		// Show already selected words
		if len(tabSelectedWord) != 0 {
			printSelected = ""
			fmt.Println(dataTranslation.WordsAlreadyUsed)
			for i := 0; i < len(tabSelectedWord); i++ {
				if i != len(tabSelectedWord)-1 {
					printSelected += tabSelectedWord[i] + " "
				} else {
					printSelected += tabSelectedWord[i]
				}
			}
			fmt.Println(printSelected)
		}

		// Draw the hangman
		DrawHangman(dataFiles, dataGame, dataTranslation, dataError, stageHangman)
		var numbOfPossibility int
		// In case the player loses the game (depending on the extension chosen the number of attempts are not the same)
		// he is asked if he wants to start a game again or not.
		if dataTranslation.VeryEasy == nameDifficulty {
			numbOfPossibility = 15
		} else {
			numbOfPossibility = 10
		}

		if stageHangman == numbOfPossibility {
			fmt.Println(dataTranslation.MessageJoseDead)
			fmt.Println(dataTranslation.MessageWordToGuess + wordsToGuess)
			for {
				fmt.Println(dataTranslation.PlayAgainQuestion + dataTranslation.Indicate + "(" + dataTranslation.PlayAgainYes + ", " + dataTranslation.PlayAgainNo + ") ;=)")
				fmt.Scanln(&tryAgain)
				tryAgain = strings.ToLower(tryAgain)
				if tryAgain == dataTranslation.PlayAgainYes {
					return dataGame.YesRetry, false
				} else if tryAgain == dataTranslation.PlayAgainNo {
					return dataGame.NoRetry, false
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}
		}

		// 'RevealRandomLettersOfWord' is used to reveal random letters of the word chosen also randomly
		// The formula is (lettersToReveal := wordLen/2 - 1)
		// If ever with the formula the count goes back to 0, there will be no revealed letter
		if stageHangman == 1 && !revealWord {
			actuelWord = RevealRandomLettersOfWord(dataTranslation, wordsToGuess, nameDifficulty)
			revealWord = true
			fmt.Println(actuelWord)
		} else {
			fmt.Println(actuelWord)
		}

		// User input request
		fmt.Println(dataTranslation.GuessLetterOrWord)
		fmt.Scanln(&selectedLetter)
		runeLetter := []rune(selectedLetter)

		// Checking the player’s input to make sure it is a single character or an error is displayed
		if len(runeLetter) == 0 {
			ClearScreen()

			manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.MissingCharacterError)

			if nameDifficulty == dataTranslation.Hacker {
				stageHangman++
			}

		} else if len(runeLetter) == 1 {
			// Checks whether the chosen character is a letter or not using the 'IsLetter' function
			// otherwise an error message will be displayed
			if !IsLetter(runeLetter[0]) {
				ClearScreen()

				manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.NotALetterError)

				if nameDifficulty == dataTranslation.Hacker {
					stageHangman++
				}
			} else if AlreadySelected(tabSelectedLetter, GetLetterConversion(runeLetter[0])) {
				ClearScreen()

				manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.AlreadySelectedLetterError)

				if nameDifficulty == dataTranslation.Hacker {
					stageHangman++
				}

			} else {
				// 1. We test a character to see if it can be converted to lowercase using the 'GetLetterConversion' function
				// 2. Then we test the presence of converted or unconverted characters in the word to be searched for using the 'TestLetterInWord' function and the 'RevealWord' function
				ClearScreen()

				letterConversion := GetLetterConversion(runeLetter[0])
				tabSelectedLetter = append(tabSelectedLetter, letterConversion)

				if TestLetterInWord(wordsToGuess, letterConversion) {
					actuelWord = RevealWord(wordsToGuess, actuelWord, letterConversion)
					if actuelWord == wordsToGuess {
						youWin = true
					}
				} else {
					stageHangman++
				}
			}

		} else if len(selectedLetter) == len(wordsToGuess) {
			// If the player enters a letter that already exists then we remove the entered letter and we make the reflection by an error message otherwise
			// or the player finds the correct letter and wins
			// either the player is increasing by one level regarding the drawing 'Hangman'
			if AlreadySelected(tabSelectedWord, selectedLetter) {
				ClearScreen()

				manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.AlreadySelectedWordError)

				if nameDifficulty == dataTranslation.Hacker {
					stageHangman++
				}
			} else {
				tabSelectedWord = append(tabSelectedWord, selectedLetter)
				if selectedLetter == wordsToGuess {
					youWin = true
				} else {
					ClearScreen()
					stageHangman++
				}
			}
		} else {
			ClearScreen()

			manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.SingleLetterOrWordError)

			if nameDifficulty == dataTranslation.Hacker {
				stageHangman++
			}
		}

		// If the player has won a win message is written otherwise it is a defeat message and then the message to start a game again is indicated
		if youWin {
			ClearScreen()
			fmt.Println(dataTranslation.Congratulations)
			fmt.Println(dataTranslation.GameWonMessage + wordsToGuess)
			for {
				fmt.Println(dataTranslation.PlayAgainQuestion + dataTranslation.Indicate + "(" + dataTranslation.PlayAgainYes + ", " + dataTranslation.PlayAgainNo + ") ;=)")
				fmt.Scanln(&tryAgain)
				tryAgain = strings.ToLower(tryAgain)
				if tryAgain == dataTranslation.PlayAgainYes {
					return dataGame.YesRetry, true
				} else if tryAgain == dataTranslation.PlayAgainNo {
					return dataGame.NoRetry, true
				} else {
					ClearScreen()
					manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
				}
			}
		}
	}
}
