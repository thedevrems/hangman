package managegame

// IsLetter allows to know the possible or not possible characters
//
// Arg :
// The only argument will be rune checked character
//
// Return :
// The function will return a boolean that will return whether or not the character is OK
func IsLetter(s rune) bool {
	if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) || (s >= 192 && s <= 197) || (s >= 199 && s <= 207) || (s >= 210 && s <= 214) || (s >= 217 && s <= 221) || (s >= 224 && s <= 229) || (s >= 232 && s <= 235) || (s >= 249 && s <= 252) {
		return true
	} else {
		return false
	}
}
