package managefiles

import (
	"fmt"
	"hangman/configuration"
	"hangman/manageerror"
	"os"
	"path/filepath"
)

func AppendWordToFile(dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, dataFiles *configuration.Structure_files, fileName string, word string) {

	filePath := filepath.Join(dataFiles.Path, fileName)
	fmt.Println(fileName, word)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.OpenFileError)
	}

	defer file.Close()

	_, err = file.WriteString(word + "\n")

	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.FileWriteError)
	}

}
