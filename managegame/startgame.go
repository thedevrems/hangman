package managegame

import (
	"fmt"
	"hangman/configuration"
	"hangman/manageerror"
	"hangman/managefiles"
	"strings"
)

// StartGame configures the start of the game with its various options and possible extensions.
//
// Args :
// Takes a pointer to a Structure_configuration object as an argument.
// Takes a pointer to a Structure_game object in the second argument
// Takes a pointer to a Structure_error object in the third argument
// Takes a pointer to a Structure_translation object in the fourth argument
// Takes a pointer to a Structure_files object in the fifth argument
//
// 1. First clear the screen
// 2. The number of Wins and the number of Losses from the start is 0 but not registered (only if the option was chosen)
// 3. Takes into account the difficulty level returned by the player (if the option is enabled) otherwise the difficulty level will be by default
// 4. The number of Wins is taken into account thanks to the generic playgame function that checks whether a win or loss has occurred 
// otherwise errors are returned by the different arguments of the function itself.
// 5. If you want to start again (thanks to the chosen extension) the screen is empty and a new game starts otherwise we stop.
func StartGame(dataConfig *configuration.Structure_configuration, dataGame *configuration.Structure_game, dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, dataFiles *configuration.Structure_files) {
	ClearScreen()

	var nbrVictory, nbrLoose int
	var selectDifficulty string

	if dataConfig.VictoryCounter {
		nbrVictory = 0
		nbrLoose = 0
	}

	if dataConfig.EnableDifficulty {
		for {
			fmt.Println(dataTranslation.DifficultyModeQuestion + " " + "( " + dataTranslation.VeryEasy + ", " + dataTranslation.Easy + ", " + dataTranslation.Middle + ", " + dataTranslation.Hard + ", " + dataTranslation.Expert + ", " + dataTranslation.Hacker + " )")
			fmt.Scanln(&selectDifficulty)
			selectDifficulty = strings.ToLower(selectDifficulty)
			if selectDifficulty == dataTranslation.VeryEasy {
				selectDifficulty = dataTranslation.VeryEasy
				ClearScreen()
				break
			} else if selectDifficulty == dataTranslation.Easy {
				selectDifficulty = dataTranslation.Easy
				ClearScreen()
				break
			} else if selectDifficulty == dataTranslation.Middle {
				selectDifficulty = dataTranslation.Middle
				ClearScreen()
				break
			} else if selectDifficulty == dataTranslation.Hard {
				selectDifficulty = dataTranslation.Hard
				ClearScreen()
				break
			} else if selectDifficulty == dataTranslation.Expert {
				selectDifficulty = dataTranslation.Expert
				ClearScreen()
				break
			} else if selectDifficulty == dataTranslation.Hacker {
				selectDifficulty = dataTranslation.Hacker
				ClearScreen()
				break
			} else {
				ClearScreen()
				manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
			}
		}
	} else {
		selectDifficulty = dataTranslation.Normal
	}

	message, isVictory := PlayHangman(dataError, dataTranslation, dataFiles, dataConfig, dataGame, 0, 0, selectDifficulty)

	if message == dataError.ErrSelectWord {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.ErrSelectWordDescription)
		return
	} else if message == dataError.ErrSelectMode {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.ErrSelectModeDescription)
		return
	}

	if dataConfig.VictoryCounter {
		if isVictory {
			nbrVictory++
		} else {
			nbrLoose++
		}
	}

	for {
		if message == dataGame.YesRetry {

			if dataConfig.AddWordAfterGame {
				ClearScreen()
				var addQuestiondWord string
				var wordToInsert string
				for {
					fmt.Println(dataTranslation.AddWordToDifficultyMode + dataTranslation.Indicate + "(" + dataTranslation.AddWordToDifficultyModeYes + ", " + dataTranslation.AddWordToDifficultyModeNo + ") ;=)")
					fmt.Scanln(&addQuestiondWord)
					addQuestiondWord = strings.ToLower(addQuestiondWord)
					if addQuestiondWord == dataTranslation.AddWordToDifficultyModeYes {
						ClearScreen()
						for {
							var correctWord bool = true
							var goodWord string
							fmt.Println(dataTranslation.InsertWord)
							fmt.Scanln(&wordToInsert)
							for _, char := range wordToInsert {
								if !IsLetter(char) {
									correctWord = false
									break
								}
							}
							if correctWord {
								var fileName string
								goodWord = wordToInsert
								if dataConfig.Language == dataGame.English {
									if selectDifficulty == dataTranslation.VeryEasy {
										fileName = dataFiles.NameFilesConfigEnglishWordsVeryEasy
									} else if selectDifficulty == dataTranslation.Easy {
										fileName = dataFiles.NameFilesConfigEnglishWordsEasy
									} else if selectDifficulty == dataTranslation.Middle {
										fileName = dataFiles.NameFilesConfigEnglishWordsMiddle
									} else if selectDifficulty == dataTranslation.Hard {
										fileName = dataFiles.NameFilesConfigEnglishWordsHard
									} else if selectDifficulty == dataTranslation.Expert {
										fileName = dataFiles.NameFilesConfigEnglishWordsExpert
									} else if selectDifficulty == dataTranslation.Hacker {
										fileName = dataFiles.NameFilesConfigEnglishWordsHacker
									} else {
										fileName = dataFiles.NameFilesConfigEnglishWordsDefault
									}
								} else {
									if selectDifficulty == dataTranslation.VeryEasy {
										fileName = dataFiles.NameFilesConfigFrenchWordsVeryEasy
									} else if selectDifficulty == dataTranslation.Easy {
										fileName = dataFiles.NameFilesConfigFrenchWordsEasy
									} else if selectDifficulty == dataTranslation.Middle {
										fileName = dataFiles.NameFilesConfigFrenchWordsMiddle
									} else if selectDifficulty == dataTranslation.Hard {
										fileName = dataFiles.NameFilesConfigFrenchWordsHard
									} else if selectDifficulty == dataTranslation.Expert {
										fileName = dataFiles.NameFilesConfigFrenchWordsExpert
									} else if selectDifficulty == dataTranslation.Hacker {
										fileName = dataFiles.NameFilesConfigFrenchWordsHacker
									} else {
										fileName = dataFiles.NameFilesConfigFrenchWordsDefault
									}
								}
								managefiles.AppendWordToFile(dataError, dataTranslation, dataFiles, fileName, goodWord)
								ClearScreen()
								break
							} else {
								ClearScreen()
								manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
							}
						}
						break
					} else if addQuestiondWord == dataTranslation.AddWordToDifficultyModeNo {
						ClearScreen()
						break
					} else {
						ClearScreen()
						manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.RespectFormulation)
					}
				}
			}

			message, isVictory = PlayHangman(dataError, dataTranslation, dataFiles, dataConfig, dataGame, nbrVictory, nbrLoose, selectDifficulty)
			if dataConfig.VictoryCounter {
				if isVictory {
					nbrVictory++
				} else {
					nbrLoose++
				}
			}

		} else if message == dataGame.NoRetry {
			ClearScreen()
			break
		} else {
			break
		}
	}
}
