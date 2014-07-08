package helpers

import (
	"os"

	"github.com/fatih/color"
)

func ExitWithError(err error) {
	color.Red("Error: %v\n", err)
	os.Exit(1)
}

func ExitWithMessage(msg string) {
	color.Red(msg)
	os.Exit(1)
}
