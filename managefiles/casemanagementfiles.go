package managefiles

import "hangman/configuration"

// CaseManagementFiles function manages the creation of necessary files for the application.
// It takes a pointer to a Structure_files object as an argument.
//
// 1. Checks if the file for configuring words (NameFilesConfigWords) exists using the CheckFiles function.
//   - If the file does not exist, it creates it using the CreateFiles function.
//
// 2. Checks if the file for displaying results (NameFilesDisplayResult) exists.
//   - If the file does not exist, it creates it as well.
//
// This ensures that the required configuration and result display files are always present for the application.
func CaseManagementFiles(dataFiles *configuration.Structure_files, dataConfiguration *configuration.Structure_configuration, dataGame *configuration.Structure_game, dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation) {
	var selectlanguage string = dataConfiguration.Language

	var englishDetected bool = selectlanguage == dataGame.English
	if englishDetected {
		if dataConfiguration.EnableDifficulty {
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsVeryEasy, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsEasy, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsMiddle, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsHard, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsExpert, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsHacker, englishDetected)
		} else {
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigEnglishWordsDefault, englishDetected)
		}
	} else {
		if dataConfiguration.EnableDifficulty {
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsVeryEasy, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsEasy, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsMiddle, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsHard, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsExpert, englishDetected)
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsHacker, englishDetected)
		} else {
			WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesConfigFrenchWordsDefault, englishDetected)
		}

	}
	WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesDisplayResult, englishDetected)
}
