package managefiles

import (
	"hangman/configuration"
	"hangman/manageerror"
	"os"
	"path/filepath"
)

// WriteFile handles the creation of a single file required by the application.
//
// Args:
// Takes a pointer to a Structure_files object as an argument.
// It also takes a pointer to a Structure_error object in order to manage possible errors
// It uses a pointer to a Structure_translation object allowing the writing of possible error
// It takes into account the path of the file to which the file is created
// It takes into account the name of the file that will be taken into account for the creation 
// The 'englishDetected' argument is also checked to see if the file is in English or French
//
// 1. First check if an error is possible at the beginning through the 'printError' function
//
// 2. Check if a .txt file can be created
//
// 3. Checks whether the word is English or not and enters it in the corresponding file
//   Generates an error if a word is already present or the word is not possible
//
// 4. Eventually write to the corresponding file 

func WriteFile(dataFiles *configuration.Structure_files, dataError *configuration.Structure_error, dataTranslation *configuration.Structure_translation, path string, nameFiles string, englishDetected bool) {
	filePath := filepath.Join(path, nameFiles)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.DirectoryCreationError+path)
		return
	}

	fileInfo, err := os.Stat(filePath)
	if err == nil && fileInfo.Size() > 0 {
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.FileCreationError+nameFiles)
		return
	}
	defer file.Close()

	var contentToWrite string

	if englishDetected {
		if dataFiles.NameFilesConfigEnglishWordsDefault == nameFiles {
			contentTemp := append(dataFiles.EnglishWordsEasy, dataFiles.EnglishWordsMiddle...)
			for _, word := range contentTemp {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigEnglishWordsVeryEasy == nameFiles {
			for _, word := range dataFiles.EnglishWordsVeryEasy {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigEnglishWordsEasy == nameFiles {
			for _, word := range dataFiles.EnglishWordsEasy {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigEnglishWordsMiddle == nameFiles {
			for _, word := range dataFiles.EnglishWordsMiddle {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigEnglishWordsHard == nameFiles {
			for _, word := range dataFiles.EnglishWordsHard {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigEnglishWordsExpert == nameFiles {
			for _, word := range dataFiles.EnglishWordsExpert {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigEnglishWordsHacker == nameFiles {
			for _, word := range dataFiles.EnglishWordsHacker {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesDisplayResult == nameFiles {
			contentToWrite = dataFiles.ContentDisplayResult
		} else {
			manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.UnsupportedFileNameError+nameFiles)
			return
		}
	} else {
		if dataFiles.NameFilesConfigFrenchWordsDefault == nameFiles {
			contentTemp := append(dataFiles.FrenchWordsEasy, dataFiles.FrenchWordsMiddle...)
			for _, word := range contentTemp {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigFrenchWordsVeryEasy == nameFiles {
			for _, word := range dataFiles.FrenchWordsVeryEasy {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigFrenchWordsEasy == nameFiles {
			for _, word := range dataFiles.FrenchWordsEasy {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigFrenchWordsMiddle == nameFiles {
			for _, word := range dataFiles.FrenchWordsMiddle {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigFrenchWordsHard == nameFiles {
			for _, word := range dataFiles.FrenchWordsHard {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigFrenchWordsExpert == nameFiles {
			for _, word := range dataFiles.FrenchWordsExpert {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesConfigFrenchWordsHacker == nameFiles {
			for _, word := range dataFiles.FrenchWordsHacker {
				contentToWrite += word + "\n"
			}
		} else if dataFiles.NameFilesDisplayResult == nameFiles {
			contentToWrite = dataFiles.ContentDisplayResult
		} else {
			manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.UnsupportedFileNameError+nameFiles)
			return
		}
	}

	if dataFiles.NameFilesDisplayResult == nameFiles {
		contentToWrite = dataFiles.ContentDisplayResult
	}

	// Write the content to the file
	_, err = file.WriteString(contentToWrite)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.FileWriteError+nameFiles)
		return
	}
}
