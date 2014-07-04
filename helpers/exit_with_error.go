package helpers

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	fmt.Printf("Error: %v", err)
	os.Exit(1)
}
