package managegame

// TestLetterInWord checks whether the letter is present in the word or not
//
// Args :
// We set the word to be found in string type
// The additional named letter written by the string player
//
// Return :
// Returns a boolean to see if the letter is present or not
func TestLetterInWord(wordsToGuess, selectedLetter string) bool {
	if len(selectedLetter) != 1 {
		return false
	}

	for _, car := range wordsToGuess {
		if string(car) == selectedLetter {
			return true
		}
	}
	return false
}
