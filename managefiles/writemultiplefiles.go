package managefiles

import "github.com/thedevrems/hangman/configuration"

// WriteMultipleFiles is a helper function to write multiple files based on difficulty levels.
func WriteMultipleFiles(dataFiles *configuration.Structure_files, dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, path string, fileNames []string, englishDetected bool) {
	for _, fileName := range fileNames {
		WriteFile(dataFiles, dataError, dataTranslation, path, fileName, englishDetected)
	}
}
