package managegame

import (
	"hangman/configuration"
	"math/rand"
	"strings"
)

func EnableJokers(dataGame *configuration.Structure_game, actualWord string, goodWord string, tabSelectedLetter []string) (word string, char string) {

	// Convertir tabSelectedLetter en map pour faciliter les vérifications
	usedLetters := make(map[string]bool)
	for _, letter := range tabSelectedLetter {
		usedLetters[letter] = true
	}

	// Trouver les lettres manquantes qui ne sont pas dans tabSelectedLetter
	var availableLetters []string
	for i, char := range goodWord {
		if actualWord[i] == '_' && !usedLetters[string(char)] {
			availableLetters = append(availableLetters, string(char))
		}
	}

	// Choisir une lettre aléatoire parmi celles disponibles
	selectedLetter := availableLetters[rand.Intn(len(availableLetters))]

	// Remplacer toutes les occurrences de la lettre choisie dans actualWord
	var newWord strings.Builder
	for i, char := range goodWord {
		if string(char) == selectedLetter || actualWord[i] != '_' {
			newWord.WriteByte(goodWord[i])
		} else {
			newWord.WriteByte('_')
		}
	}

	// Vérifier si le nouveau mot est complet
	finalWord := newWord.String()
	if finalWord == goodWord {
		return dataGame.ErrorJokers, selectedLetter
	}

	// Retourner le nouveau mot et la lettre révélée
	return finalWord, selectedLetter
}
