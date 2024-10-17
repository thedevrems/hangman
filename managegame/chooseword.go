package managegame

import (
	"hangman/configuration"
	"math/rand"
)

// 'ChooseWord' selects a random word, converts it, and checks if all words are valid.

// Args:
// Takes a pointer to a Structure_error object as an argument.
// Takes a pointer to a Structure_translation object in the second argument
// Takes a pointer to a Structure_files object in the third argument
// Takes a pointer to a Structure_configuration object in the fourth argument
// Takes a pointer to a Structure_game object in the fifth argument
// And the difficulty chosen by the player
//
// 1. If a word contains an invalid character, it will move to another word without repeating the words already tested.
//
// 2. Try every word until a valid word is found or everything is tested
//   Choose a random word
//   Check if the index has already been used
//   Switch to another if it has already been tested
//
// 3. Check if the index is already in use
//
// 4. Checks if the letter is correct and converts each letter of the word with the function 'GetLetterConversion'
//   Checks if the converted letter is valid with the function 'IsLetter'
//
// 5. If one letter is invalid, the word is rejected and another
//   If the word is valid, it is returned

func ChooseWord(dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, dataFiles *configuration.Structure_files, dataConfig *configuration.Structure_configuration, dataGame *configuration.Structure_game, nameDifficulty string) (word string) {
	var tabString []string
	var errSelectMode string = dataError.ErrSelectMode
	var errSelectWord string = dataError.ErrSelectWord

	if dataConfig.Language == dataGame.English {
		if nameDifficulty == dataTranslation.VeryEasy {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsVeryEasy)
		} else if nameDifficulty == dataTranslation.Easy {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsEasy)
		} else if nameDifficulty == dataTranslation.Middle {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsMiddle)
		} else if nameDifficulty == dataTranslation.Hard {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsHard)
		} else if nameDifficulty == dataTranslation.Expert {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsExpert)
		} else if nameDifficulty == dataTranslation.Hacker {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsHacker)
		} else if nameDifficulty == dataTranslation.Normal {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsDefault)
		} else {
			return errSelectMode
		}
	} else {
		if nameDifficulty == dataTranslation.VeryEasy {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsVeryEasy)
		} else if nameDifficulty == dataTranslation.Easy {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsEasy)
		} else if nameDifficulty == dataTranslation.Middle {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsMiddle)
		} else if nameDifficulty == dataTranslation.Hard {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsHard)
		} else if nameDifficulty == dataTranslation.Expert {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsExpert)
		} else if nameDifficulty == dataTranslation.Hacker {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsHacker)
		} else if nameDifficulty == dataTranslation.Normal {
			tabString = ReadFile(dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsDefault)
		} else {
			return errSelectMode
		}
	}

	usedIndices := make(map[int]bool)

	for len(usedIndices) < len(tabString) {
		randomIndex := rand.Intn(len(tabString))

		if usedIndices[randomIndex] {
			continue
		}

		usedIndices[randomIndex] = true

		wrongWord := tabString[randomIndex]

		if wrongWord == errSelectMode || wrongWord == errSelectWord || len(wrongWord) < 2 {
			continue
		}

		goodWord := ""

		isValidWord := true
		for _, v := range wrongWord {
			goodLetter := GetLetterConversion(v)
			goodLetterRune := []rune(goodLetter)

			if len(goodLetterRune) > 0 && IsLetter(goodLetterRune[0]) {
				goodWord += string(goodLetterRune)
			} else {
				isValidWord = false
				break
			}
		}

		if isValidWord {
			return goodWord
		}
	}

	return errSelectWord
}
