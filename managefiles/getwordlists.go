package managefiles

import "hangman/configuration"

// getWordLists returns the appropriate word list based on the file name and language
func getWordLists(dataFiles *configuration.Structure_files, nameFiles string, englishDetected bool) []string {
	if englishDetected {
		switch nameFiles {
		case dataFiles.NameFilesConfigEnglishWordsDefault:
			return append(dataFiles.EnglishWordsEasy, dataFiles.EnglishWordsMiddle...)
		case dataFiles.NameFilesConfigEnglishWordsVeryEasy:
			return dataFiles.EnglishWordsVeryEasy
		case dataFiles.NameFilesConfigEnglishWordsEasy:
			return dataFiles.EnglishWordsEasy
		case dataFiles.NameFilesConfigEnglishWordsMiddle:
			return dataFiles.EnglishWordsMiddle
		case dataFiles.NameFilesConfigEnglishWordsHard:
			return dataFiles.EnglishWordsHard
		case dataFiles.NameFilesConfigEnglishWordsExpert:
			return dataFiles.EnglishWordsExpert
		case dataFiles.NameFilesConfigEnglishWordsHacker:
			return dataFiles.EnglishWordsHacker
		case dataFiles.NameFilesDisplayResult:
			return []string{dataFiles.ContentDisplayResult}
		default:
			return nil
		}
	} else {
		switch nameFiles {
		case dataFiles.NameFilesConfigFrenchWordsDefault:
			return append(dataFiles.FrenchWordsEasy, dataFiles.FrenchWordsMiddle...)
		case dataFiles.NameFilesConfigFrenchWordsVeryEasy:
			return dataFiles.FrenchWordsVeryEasy
		case dataFiles.NameFilesConfigFrenchWordsEasy:
			return dataFiles.FrenchWordsEasy
		case dataFiles.NameFilesConfigFrenchWordsMiddle:
			return dataFiles.FrenchWordsMiddle
		case dataFiles.NameFilesConfigFrenchWordsHard:
			return dataFiles.FrenchWordsHard
		case dataFiles.NameFilesConfigFrenchWordsExpert:
			return dataFiles.FrenchWordsExpert
		case dataFiles.NameFilesConfigFrenchWordsHacker:
			return dataFiles.FrenchWordsHacker
		case dataFiles.NameFilesDisplayResult:
			return []string{dataFiles.ContentDisplayResult}
		default:
			return nil
		}
	}
}
