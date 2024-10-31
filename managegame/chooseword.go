package managegame

import (
	"math/rand"

	"github.com/thedevrems/hangman/configuration"
)

// ChooseWord selects a random word, converts it, and checks if it's valid.
func ChooseWord(dataError *configuration.Structure_error, dataTranslation *configuration.TranslationHangman, dataFiles *configuration.Structure_files, dataConfig *configuration.Structure_configuration, dataGame *configuration.Structure_game, nameDifficulty string) (word string) {
	var errSelectMode = dataError.ErrSelectMode
	var errSelectWord = dataError.ErrSelectWord

	// Mapping between difficulty levels and file names
	difficultyFiles := map[string]string{
		dataTranslation.VeryEasy: dataFiles.NameFilesConfigEnglishWordsVeryEasy,
		dataTranslation.Easy:     dataFiles.NameFilesConfigEnglishWordsEasy,
		dataTranslation.Middle:   dataFiles.NameFilesConfigEnglishWordsMiddle,
		dataTranslation.Hard:     dataFiles.NameFilesConfigEnglishWordsHard,
		dataTranslation.Expert:   dataFiles.NameFilesConfigEnglishWordsExpert,
		dataTranslation.Hacker:   dataFiles.NameFilesConfigEnglishWordsHacker,
		dataTranslation.Normal:   dataFiles.NameFilesConfigEnglishWordsDefault,
	}

	// Select the correct language files
	var wordList []string
	if dataConfig.Language == dataGame.English {
		wordList = ReadFile(dataError, dataTranslation, dataFiles.Path, difficultyFiles[nameDifficulty])
	} else {
		// Replace the English filenames by French ones
		difficultyFiles[dataTranslation.VeryEasy] = dataFiles.NameFilesConfigFrenchWordsVeryEasy
		difficultyFiles[dataTranslation.Easy] = dataFiles.NameFilesConfigFrenchWordsEasy
		difficultyFiles[dataTranslation.Middle] = dataFiles.NameFilesConfigFrenchWordsMiddle
		difficultyFiles[dataTranslation.Hard] = dataFiles.NameFilesConfigFrenchWordsHard
		difficultyFiles[dataTranslation.Expert] = dataFiles.NameFilesConfigFrenchWordsExpert
		difficultyFiles[dataTranslation.Hacker] = dataFiles.NameFilesConfigFrenchWordsHacker
		difficultyFiles[dataTranslation.Normal] = dataFiles.NameFilesConfigFrenchWordsDefault

		wordList = ReadFile(dataError, dataTranslation, dataFiles.Path, difficultyFiles[nameDifficulty])
	}

	// If the difficulty level is not found, return an error
	if len(wordList) == 0 {
		return errSelectMode
	}

	// Track used indices to avoid repeating words
	usedIndices := make(map[int]bool)

	for len(usedIndices) < len(wordList) {
		randomIndex := rand.Intn(len(wordList))
		if usedIndices[randomIndex] {
			continue
		}
		usedIndices[randomIndex] = true

		selectedWord := wordList[randomIndex]

		// Check if the word is invalid or too short
		if selectedWord == errSelectMode || selectedWord == errSelectWord || len(selectedWord) < 2 {
			continue
		}

		// Validate and convert the word letter by letter
		validWord := ""
		isValid := true
		for _, char := range selectedWord {
			convertedLetter := GetLetterConversion(char)
			if len(convertedLetter) > 0 && IsLetter([]rune(convertedLetter)[0]) {
				validWord += convertedLetter
			} else {
				isValid = false
				break
			}
		}

		if isValid {
			return validWord
		}
	}

	return errSelectWord
}
