package managegame

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the entire window to run a new command
func ClearScreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
