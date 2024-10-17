package managegame

import (
	"bufio"
	"hangman/configuration"
	"hangman/manageerror"
	"os"
	"path/filepath"
)

// ReadFile to read the correct file requested by the player regarding the difficulty level
// If the file cannot be opened, errors will be returned.
// The bufio package is used which provides read and write classes for data streams.
// It allows efficient management of I/O using buffers to reduce the number of system calls.
//
// Args :
// Takes a pointer to a Structure_error object as an argument.
// Takes a pointer to a Structure_translation object in the second argument.
// Takes the path used by the file in question with a string type
// Takes the file name of course as argument under string type
//
// Return :
// The function will return a word array corresponding to all of the words in the file
func ReadFile(dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, path string, nameFile string) []string {
	filePath := filepath.Join(path, nameFile)

	file, err := os.Open(filePath)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.OpenFileError+nameFile)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.ScanFileError+nameFile)
	}

	return lines
}
