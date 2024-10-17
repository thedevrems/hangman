package managegame

// GetLetterConversion allows to manage all UTF-8 characters
// Thus transform all special characters or Uppercase characters into a single lowercase character
//
// Arg :
// The only argument will be rune checked character
//
// Return :
// The function will return a string that will match the corresponding lowercase letter
func GetLetterConversion(let rune) string {
	if let >= 65 && let <= 90 { // rune(65) = 'A' & rune(90) = 'Z'
		let = let + 32 // Convert uppercase to lowercase
		return string(let)
	} else if let >= 97 && let <= 122 { // rune(97) = 'a' & rune(122) = 'z'
		return string(let)
	} else if (let >= 192 && let <= 197) || (let >= 224 && let <= 229) { // À-Å or à-å
		return "a"
	} else if (let >= 200 && let <= 203) || (let >= 232 && let <= 235) { // È-Ë or è-ë
		return "e"
	} else if let == 199 || let == 231 { // Ç or ç
		return "c"
	} else if (let >= 204 && let <= 207) || (let >= 236 && let <= 239) { // Ì-Ï or ì-ï
		return "i"
	} else if (let >= 210 && let <= 214) || (let >= 242 && let <= 246) { // Ò-Ö or ò-ö
		return "o"
	} else if (let >= 217 && let <= 220) || (let >= 249 && let <= 252) { // Ù-Ü or ù-ü
		return "u"
	} else {
		return ""
	}
}
