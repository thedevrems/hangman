package configuration

// 'ConfigError' initializes and returns a pointer to a Structure_error struct !

// Return :
// The different colors involved in a possible error during the game phase

func ConfigError() *Structure_error {
	return &Structure_error{
		Reset:         "\033[0m",
		Red:           "\033[31m",
		Green:         "\033[32m",
		Yellow:        "\033[33m",
		Blue:          "\033[34m",
		Purple:        "\033[35m",
		Cyan:          "\033[36m",
		White:         "\033[37m",
		ErrSelectMode: "errorSelectMode",
		ErrSelectWord: "errorSelectWord",
	}
}
