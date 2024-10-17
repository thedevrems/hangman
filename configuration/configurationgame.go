package configuration

// ConfigGame initializes and returns a pointer to a Structure_game struct

// Return
// The drawing in the hangman.txt file corresponding to the number of wrong(s) answer(s) he made
func ConfigGame() *Structure_game {
	return &Structure_game{
		DrawHangmanLineStartPos1:  3,
		DrawHangmanLineStopPos1:   3,
		DrawHangmanLineStartPos2:  7,
		DrawHangmanLineStopPos2:   12,
		DrawHangmanLineStartPos3:  16,
		DrawHangmanLineStopPos3:   22,
		DrawHangmanLineStartPos4:  26,
		DrawHangmanLineStopPos4:   32,
		DrawHangmanLineStartPos5:  36,
		DrawHangmanLineStopPos5:   42,
		DrawHangmanLineStartPos6:  46,
		DrawHangmanLineStopPos6:   52,
		DrawHangmanLineStartPos7:  56,
		DrawHangmanLineStopPos7:   62,
		DrawHangmanLineStartPos8:  66,
		DrawHangmanLineStopPos8:   72,
		DrawHangmanLineStartPos9:  76,
		DrawHangmanLineStopPos9:   82,
		DrawHangmanLineStartPos10: 86,
		DrawHangmanLineStopPos10:  92,
		DrawHangmanLineStartPos11: 96,
		DrawHangmanLineStopPos11:  103,
		DrawHangmanLineStartPos12: 107,
		DrawHangmanLineStopPos12:  115,
		DrawHangmanLineStartPos13: 119,
		DrawHangmanLineStopPos13:  127,
		DrawHangmanLineStartPos14: 131,
		DrawHangmanLineStopPos14:  139,
		DrawHangmanLineStartPos15: 143,
		DrawHangmanLineStopPos15:  151,
		English:                   "en",
		Frensh:                    "fr",
		YesRetry:                  "yesRetry",
		NoRetry:                   "NoRetry",
		
	}
}
