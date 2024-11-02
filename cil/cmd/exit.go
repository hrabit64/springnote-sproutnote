package cmd

import (
	"fmt"
	"os"
)

func RunExit() {
	fmt.Println("Exiting...")
	os.Exit(0)
}
