package manageerror

import (
	"fmt"

	"github.com/thedevrems/hangman/configuration"
)

// PrintError allows to display an error based on it

// Args :
// Take as first argument a pointer to a DataError object
// Take into account 'title' of the string type error
// Consider typeOfError of string type that speaks for itself

func PrintError(dataError *configuration.DataError, title string, typeOfError string) {
	fmt.Println(dataError.Red, title, dataError.Reset, typeOfError)
}
