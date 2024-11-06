package managefiles

import (
	"os"
	"path/filepath"

	"github.com/thedevrems/hangman/manageerror"

	"github.com/thedevrems/hangman/configuration"
)

// WriteFile handles the creation of a single file required by the application.
func WriteFile(dataFiles *configuration.DataFiles, dataError *configuration.DataError, dataTranslation *configuration.TranslationHangman, path string, nameFiles string, englishDetected bool) {
	filePath := filepath.Join(path, nameFiles)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.DirectoryCreationError+path)
		return
	}

	// Check if the file exists and has content
	fileInfo, err := os.Stat(filePath)
	if err == nil && fileInfo.Size() > 0 {
		return
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.FileCreationError+nameFiles)
		return
	}
	defer file.Close()

	// Select the content to write based on language and file name
	var contentToWrite string
	wordLists := getWordLists(dataFiles, nameFiles, englishDetected)

	if wordLists == nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.UnsupportedFileNameError+nameFiles)
		return
	}

	// Join the word lists into a single string
	for _, word := range wordLists {
		contentToWrite += word + "\n"
	}

	// Write content to the file
	_, err = file.WriteString(contentToWrite)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.FileWriteError+nameFiles)
		return
	}
}
