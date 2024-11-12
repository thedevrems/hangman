package managefiles

import "github.com/thedevrems/hangman/configuration"

// WriteMultipleFiles is a helper function to write multiple files based on difficulty levels.
func WriteMultipleFiles(dataFiles *configuration.DataFiles, dataError *configuration.DataError, dataTranslation *configuration.TranslationHangman, path string, fileNames []string, englishDetected bool) {
	for _, fileName := range fileNames {
		WriteFile(dataFiles, dataError, dataTranslation, path, fileName, englishDetected)
	}
}
