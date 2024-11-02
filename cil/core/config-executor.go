package core

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/cmd"
)

func processConfigExecutor(args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
		return
	}

	switch args[1] {
	case "show":
		cmd.RunShowConfig()
		return
	case "edit":
		if len(args) < 3 {
			fmt.Println("Invalid command. Type 'help config' for a list of commands.")
			return
		} else if len(args) < 4 {
			fmt.Println("Invalid command. Type 'help config' for a list of commands.")
			return
		}
		cmd.RunConfigEdit(args[2], args[3])
		return
	}

}
