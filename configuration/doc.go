package configuration

// Corresponds to struct DataError structure
type DataError struct {
	Reset         string
	Red           string
	Green         string
	Yellow        string
	Blue          string
	Purple        string
	Cyan          string
	White         string
	ErrSelectMode string
	ErrSelectWord string
}

// Corresponds to struct GameHangman structure which returns the position of the line in the Hangman.txt file
type DataGameHangman struct {
	DrawHangmanLineStartPos1  int
	DrawHangmanLineStopPos1   int
	DrawHangmanLineStartPos2  int
	DrawHangmanLineStopPos2   int
	DrawHangmanLineStartPos3  int
	DrawHangmanLineStopPos3   int
	DrawHangmanLineStartPos4  int
	DrawHangmanLineStopPos4   int
	DrawHangmanLineStartPos5  int
	DrawHangmanLineStopPos5   int
	DrawHangmanLineStartPos6  int
	DrawHangmanLineStopPos6   int
	DrawHangmanLineStartPos7  int
	DrawHangmanLineStopPos7   int
	DrawHangmanLineStartPos8  int
	DrawHangmanLineStopPos8   int
	DrawHangmanLineStartPos9  int
	DrawHangmanLineStopPos9   int
	DrawHangmanLineStartPos10 int
	DrawHangmanLineStopPos10  int
	DrawHangmanLineStartPos11 int
	DrawHangmanLineStopPos11  int
	DrawHangmanLineStartPos12 int
	DrawHangmanLineStopPos12  int
	DrawHangmanLineStartPos13 int
	DrawHangmanLineStopPos13  int
	DrawHangmanLineStartPos14 int
	DrawHangmanLineStopPos14  int
	DrawHangmanLineStartPos15 int
	DrawHangmanLineStopPos15  int
	English                   string
	Frensh                    string
	YesRetry                  string
	NoRetry                   string
	ErrorJokers               string
	NumberModJokers           rune
}

// Corresponds to struct DataFiles structure
type DataFiles struct {
	NameFilesConfigFrenchWordsDefault  string
	NameFilesConfigFrenchWordsVeryEasy string
	NameFilesConfigFrenchWordsEasy     string
	NameFilesConfigFrenchWordsMiddle   string
	NameFilesConfigFrenchWordsHard     string
	NameFilesConfigFrenchWordsExpert   string
	NameFilesConfigFrenchWordsHacker   string

	NameFilesConfigEnglishWordsDefault  string
	NameFilesConfigEnglishWordsVeryEasy string
	NameFilesConfigEnglishWordsEasy     string
	NameFilesConfigEnglishWordsMiddle   string
	NameFilesConfigEnglishWordsHard     string
	NameFilesConfigEnglishWordsExpert   string
	NameFilesConfigEnglishWordsHacker   string

	NameFilesDisplayResult string
	Path                   string

	FrenchWordsVeryEasy []string
	FrenchWordsEasy     []string
	FrenchWordsMiddle   []string
	FrenchWordsHard     []string
	FrenchWordsExpert   []string
	FrenchWordsHacker   []string

	EnglishWordsVeryEasy []string
	EnglishWordsEasy     []string
	EnglishWordsMiddle   []string
	EnglishWordsHard     []string
	EnglishWordsExpert   []string
	EnglishWordsHacker   []string

	ContentDisplayResult string
}

// Corresponds to struct DataConfiguration structure
type DataConfigHangman struct {
	Language         string
	VictoryCounter   bool
	EnableDifficulty bool
	AddWordAfterGame bool
	EnableJokers     bool
}

// Corresponds to struct TranslationHangman structure
type TranslationHangman struct {
	UnsupportedFileNameError   string
	DirectoryCreationError     string
	FileCreationError          string
	FileWriteError             string
	TittleError                string
	DifficultyModeQuestion     string
	VeryEasy                   string
	Easy                       string
	Middle                     string
	Hard                       string
	Expert                     string
	Hacker                     string
	Normal                     string
	RespectFormulation         string
	ErrSelectWordDescription   string
	ErrSelectModeDescription   string
	OpenFileError              string
	ScanFileError              string
	InvalidPositionHangman     string
	TitleHangman               string
	YourCurrentlyOn            string
	Victory                    string
	Defeat                     string
	And                        string
	LettersAlreadyUsed         string
	WordsAlreadyUsed           string
	MessageJoseDead            string
	MessageWordToGuess         string
	PlayAgainQuestion          string
	Indicate                   string
	PlayAgainYes               string
	PlayAgainNo                string
	GuessLetterOrWord          string
	MissingCharacterError      string
	NotALetterError            string
	AlreadySelectedLetterError string
	AlreadySelectedWordError   string
	SingleLetterOrWordError    string
	Congratulations            string
	GameWonMessage             string
	AddWordToDifficultyMode    string
	AddWordToDifficultyModeYes string
	AddWordToDifficultyModeNo  string
	SelectModYes               string
	SelectModNo                string
	InsertWord                 string
	Enable                     string
	Disable                    string
	ConfigurationCurrent       string
	EnterNumber                string
	OptionLanguage             string
	OptionVictoryCounter       string
	OptionEnableDifficulty     string
	OptionAddWordAfterGame     string
	OptionStartGame            string
	AllowAddingWordAtEnd       string
	EnableWinCounter           string
	EnableDifficulty           string
	PleaseEnterLanguage        string
	LeaveTheGame               string
	OptionEnableJokers         string
	EnableJokers               string
	InsertJokerMessage         string
	JokersCountMessage         string
	JokerLimitMessage          string
	JokerUsedMessage           string
}
