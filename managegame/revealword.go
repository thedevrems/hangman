package managegame

// RevealWord is a function that allows you to reveal a word
//
// Args :
// As the first argument we indicate the words to be found of type string
// In the second argument we indicate the currently chosen word of type string
// As the last argument we indicate the selected letter of type string
//
// Return :
// Returns a string

func RevealWord(wordsToGuess string, actuelWord string, selectedLetter string) string {
	runesToGuess := []rune(wordsToGuess)
	runesActuel := []rune(actuelWord)
	runeLetter := []rune(selectedLetter)[0]

	for i, r := range runesToGuess {
		if r == runeLetter {
			runesActuel[i] = r
		}
	}

	return string(runesActuel)
}
