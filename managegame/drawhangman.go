package managegame

import (
	"bufio"
	"fmt"
	"hangman/configuration"
	"hangman/manageerror"
	"os"
	"path/filepath"
	"strings"
)

// DrawHangman will draw the hanged all along the player’s chess
// We use the switch which is much more logical in this kind of case
// So at each loss we increase the case and therefore in the 'switch' the 'case' with the start line and the end line regarding the drawing of the Hangman
// If there is a positioning error, an error is displayed.
// The first line of course starts at line 1.
// Then, each time the drawing is erased and replaced by the next one.
//
// Args :
// Takes a pointer to a Structure_files object as an argument.
// Takes a pointer to a Structure_game object in the second argument
// Takes a pointer to a Structure_translation object in the third argument
// Takes a pointer to a Structure_error object in the fourth argument
// And the communication of the position of the type drawing int

func DrawHangman(dataFiles *configuration.Structure_files, dataGame *configuration.Structure_game, dataTranslation *configuration.Structure_translation, dataError *configuration.Structure_error, pos int) {
	var lineStart int
	var lineStop int

	switch pos {
	case 1:
		lineStart = dataGame.DrawHangmanLineStartPos1
		lineStop = dataGame.DrawHangmanLineStopPos1
	case 2:
		lineStart = dataGame.DrawHangmanLineStartPos2
		lineStop = dataGame.DrawHangmanLineStopPos2
	case 3:
		lineStart = dataGame.DrawHangmanLineStartPos3
		lineStop = dataGame.DrawHangmanLineStopPos3
	case 4:
		lineStart = dataGame.DrawHangmanLineStartPos4
		lineStop = dataGame.DrawHangmanLineStopPos4
	case 5:
		lineStart = dataGame.DrawHangmanLineStartPos5
		lineStop = dataGame.DrawHangmanLineStopPos5
	case 6:
		lineStart = dataGame.DrawHangmanLineStartPos6
		lineStop = dataGame.DrawHangmanLineStopPos6
	case 7:
		lineStart = dataGame.DrawHangmanLineStartPos7
		lineStop = dataGame.DrawHangmanLineStopPos7
	case 8:
		lineStart = dataGame.DrawHangmanLineStartPos8
		lineStop = dataGame.DrawHangmanLineStopPos8
	case 9:
		lineStart = dataGame.DrawHangmanLineStartPos9
		lineStop = dataGame.DrawHangmanLineStopPos9
	case 10:
		lineStart = dataGame.DrawHangmanLineStartPos10
		lineStop = dataGame.DrawHangmanLineStopPos10
	case 11:
		lineStart = dataGame.DrawHangmanLineStartPos11
		lineStop = dataGame.DrawHangmanLineStopPos11
	case 12:
		lineStart = dataGame.DrawHangmanLineStartPos12
		lineStop = dataGame.DrawHangmanLineStopPos12
	case 13:
		lineStart = dataGame.DrawHangmanLineStartPos13
		lineStop = dataGame.DrawHangmanLineStopPos13
	case 14:
		lineStart = dataGame.DrawHangmanLineStartPos14
		lineStop = dataGame.DrawHangmanLineStopPos14
	case 15:
		lineStart = dataGame.DrawHangmanLineStartPos15
		lineStop = dataGame.DrawHangmanLineStopPos15
	default:
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.InvalidPositionHangman)
		return
	}

	filePath := filepath.Join(dataFiles.Path, dataFiles.NameFilesDisplayResult)

	file, err := os.Open(filePath)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.OpenFileError+dataFiles.NameFilesDisplayResult)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1

	for scanner.Scan() {
		if currentLine >= lineStart && currentLine <= lineStop {
			line := strings.ReplaceAll(scanner.Text(), "\t", "    ")
			fmt.Println(line)
		}
		currentLine++
	}

	fmt.Println()

}
