package helpers

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func ExitWithMessage(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
