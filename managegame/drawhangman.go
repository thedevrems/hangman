package managegame

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thedevrems/hangman/manageerror"

	"github.com/thedevrems/hangman/configuration"
)

// DrawHangman draws the hanged man along the player's chess game.
// Each time the player loses, a new part of the hangman is drawn by updating the lines based on the given position.
// Args:
// - Takes pointers to various configuration structures (dataFiles, dataGame, etc.)
// - Takes the position 'pos' (an integer) to indicate the part of the hangman to draw.
func DrawHangman(dataFiles *configuration.Structure_files, dataGame *configuration.Structure_game, dataTranslation *configuration.Structure_translation, dataError *configuration.Structure_error, pos int) {
	// Array to hold start and stop positions for each drawing
	drawPositions := [][2]int{
		{dataGame.DrawHangmanLineStartPos1, dataGame.DrawHangmanLineStopPos1},
		{dataGame.DrawHangmanLineStartPos2, dataGame.DrawHangmanLineStopPos2},
		{dataGame.DrawHangmanLineStartPos3, dataGame.DrawHangmanLineStopPos3},
		{dataGame.DrawHangmanLineStartPos4, dataGame.DrawHangmanLineStopPos4},
		{dataGame.DrawHangmanLineStartPos5, dataGame.DrawHangmanLineStopPos5},
		{dataGame.DrawHangmanLineStartPos6, dataGame.DrawHangmanLineStopPos6},
		{dataGame.DrawHangmanLineStartPos7, dataGame.DrawHangmanLineStopPos7},
		{dataGame.DrawHangmanLineStartPos8, dataGame.DrawHangmanLineStopPos8},
		{dataGame.DrawHangmanLineStartPos9, dataGame.DrawHangmanLineStopPos9},
		{dataGame.DrawHangmanLineStartPos10, dataGame.DrawHangmanLineStopPos10},
		{dataGame.DrawHangmanLineStartPos11, dataGame.DrawHangmanLineStopPos11},
		{dataGame.DrawHangmanLineStartPos12, dataGame.DrawHangmanLineStopPos12},
		{dataGame.DrawHangmanLineStartPos13, dataGame.DrawHangmanLineStopPos13},
		{dataGame.DrawHangmanLineStartPos14, dataGame.DrawHangmanLineStopPos14},
		{dataGame.DrawHangmanLineStartPos15, dataGame.DrawHangmanLineStopPos15},
	}

	// Check if the provided position is valid
	if pos < 1 || pos > len(drawPositions) {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.InvalidPositionHangman)
		return
	}

	// Get the start and stop lines for the current position
	lineStart, lineStop := drawPositions[pos-1][0], drawPositions[pos-1][1]

	// Open the file that contains the hangman drawings
	filePath := filepath.Join(dataFiles.Path, dataFiles.NameFilesDisplayResult)
	file, err := os.Open(filePath)
	if err != nil {
		manageerror.PrintError(dataError, dataTranslation.TittleError, dataTranslation.OpenFileError+dataFiles.NameFilesDisplayResult)
		return
	}
	defer file.Close()

	// Read and display the drawing lines from the file
	scanner := bufio.NewScanner(file)
	for currentLine := 1; scanner.Scan(); currentLine++ {
		if currentLine >= lineStart && currentLine <= lineStop {
			// Replace tabs with spaces for consistent formatting
			line := strings.ReplaceAll(scanner.Text(), "\t", "    ")
			fmt.Println(line)
		}
	}

	fmt.Println()
}
