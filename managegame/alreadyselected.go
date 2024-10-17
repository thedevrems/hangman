package managegame

// AlreadySelected allows to express if a letter has already been selected or not
//
// Args :
// We will take into account the array of characters already selected with a string table type
// The second argument will display a list of letters under string type
//
// Return :
// Of course this function will return a boolean to know if the letter has already been used by the player
func AlreadySelected(selectedLetters []string, list string) bool {
	for _, l := range selectedLetters {
		if l == list {
			return true
		}
	}
	return false
}
