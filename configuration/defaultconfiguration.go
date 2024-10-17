package configuration

// DefaultConfiguration initializes and returns a pointer to a Structure_configuration struct.
//
// Return:
// The 'Language' is set to English ("en") by default, but can be changed by the player.
// The 'VictoryCounter' is initially set to false, meaning no win/loss counter is active.
// The 'EnableDifficulty' is set to false, so difficulty adjustment is not enabled by default.
// The 'AddWordAfterGame' is also set to false, meaning no word will be added after each game.
//
// The player can modify these settings later in the game panel.
func DefaultConfiguration() *Structure_configuration {
	return &Structure_configuration{
		Language: "en", // "fr" or "en"
		// Mode
		VictoryCounter:   false,
		EnableDifficulty: false,
		AddWordAfterGame: false,
	}
}
