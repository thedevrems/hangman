package managefiles

import "github.com/thedevrems/hangman/configuration"

// CaseManagementFiles manages the creation of necessary files for the application.
func CaseManagementFiles(dataFiles *configuration.Structure_files, dataConfiguration *configuration.Structure_configuration, dataGame *configuration.Structure_game, dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation) {
	var englishDetected bool = dataConfiguration.Language == dataGame.English
	var fileNames []string

	// Determine which set of files to write based on the language and difficulty.
	if englishDetected {
		if dataConfiguration.EnableDifficulty {
			fileNames = []string{
				dataFiles.NameFilesConfigEnglishWordsVeryEasy,
				dataFiles.NameFilesConfigEnglishWordsEasy,
				dataFiles.NameFilesConfigEnglishWordsMiddle,
				dataFiles.NameFilesConfigEnglishWordsHard,
				dataFiles.NameFilesConfigEnglishWordsExpert,
				dataFiles.NameFilesConfigEnglishWordsHacker,
			}
		} else {
			fileNames = []string{dataFiles.NameFilesConfigEnglishWordsDefault}
		}
	} else {
		if dataConfiguration.EnableDifficulty {
			fileNames = []string{
				dataFiles.NameFilesConfigFrenchWordsVeryEasy,
				dataFiles.NameFilesConfigFrenchWordsEasy,
				dataFiles.NameFilesConfigFrenchWordsMiddle,
				dataFiles.NameFilesConfigFrenchWordsHard,
				dataFiles.NameFilesConfigFrenchWordsExpert,
				dataFiles.NameFilesConfigFrenchWordsHacker,
			}
		} else {
			fileNames = []string{dataFiles.NameFilesConfigFrenchWordsDefault}
		}
	}

	// Write the necessary word files based on the determined names.
	WriteMultipleFiles(dataFiles, dataError, dataTranslation, dataFiles.Path, fileNames, englishDetected)

	// Always write the display result file.
	WriteFile(dataFiles, dataError, dataTranslation, dataFiles.Path, dataFiles.NameFilesDisplayResult, englishDetected)
}
