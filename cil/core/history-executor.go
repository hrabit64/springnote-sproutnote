package core

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/cmd"
	"strconv"
)

func processHistoryExecutor(args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
		return
	}

	switch args[1] {
	case "db":
		progressDBHistoryShow(args)
		return

	case "file":
		progressFileHistoryShow(args)
		return
	}

}

func progressFileHistoryShow(args []string) bool {
	if len(args) < 3 {
		fmt.Println("Invalid command. Type 'help db' for a list of commands.")
		return true
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid id.")
		return true
	}

	cmd.RunHistoryFileItemShow(id)
	return false
}

func progressDBHistoryShow(args []string) {
	if len(args) < 3 {
		fmt.Println("Invalid command. Type 'help db' for a list of commands.")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid id.")
		return
	}

	cmd.RunHistoryDBItemShow(id)
}
