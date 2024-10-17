package managegame

import (
	"hangman/configuration"
	"math/rand"
)

// RevealRandomLettersOfWord is used to reveal random letters of the word chosen also randomly
// The formula is (lettersToReveal := wordLen/2 - 1)
// If ever with the formula the count goes back to 0, there will be no revealed letter
//
// Args :
// We argument the word that was chosen from type string
//
// Return :
// The function will return the revealed strings
func RevealRandomLettersOfWord(dataTranslation *configuration.Structure_translation, word string, nameDifficulty string) string {
	wordLen := len(word)

	if nameDifficulty == dataTranslation.Hacker {
		var hardWord string
		var count int = 0

		for count < wordLen {
			hardWord += "_"
			count++
		}

		return hardWord

	} else {
		lettersToReveal := wordLen/2 - 1

		revealed := make([]rune, wordLen)

		for i := range revealed {
			revealed[i] = '_'
		}

		revealedCount := 0
		for revealedCount < lettersToReveal {
			randomIndex := rand.Intn(wordLen)

			if revealed[randomIndex] == '_' {
				revealed[randomIndex] = rune(word[randomIndex])
				revealedCount++
			}
		}

		return string(revealed)
	}

}
