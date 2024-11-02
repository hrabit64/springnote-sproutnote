package core

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/cmd"
	"strconv"
)

func processDbExecutor(args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
		return
	}

	switch args[1] {

	case "add":
		cmd.RunDBCreate()
		return

	case "show":
		progressRunDBShow(args)
		return

	case "dump":
		progressDBDump(args)
		return

	case "remove":
		progressDBRemove(args)
		return

	default:
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
	}
}

func progressDBRemove(args []string) {
	if len(args) < 3 {
		fmt.Println("Missing id.")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid id.")
		return
	}
	if id < 1 {
		fmt.Println("Invalid target number.")
		return
	}
	cmd.RunDBRemove(id)
	return
}

func progressDBDump(args []string) {
	if len(args) < 3 {
		cmd.RunDBDumpAll()
		return
	} else {
		target, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid target number.")
			return
		}
		if target < 1 {
			fmt.Println("Invalid target number.")
			return
		}
		cmd.RunDBDumpSingle(target)
		return
	}
}

func progressRunDBShow(args []string) {
	page := 1
	if len(args) >= 3 {
		page, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid page number.")
			return
		}
		if page < 1 {
			fmt.Println("Invalid page number.")
			return
		}
	}
	cmd.RunDBShow(page - 1)
	return
}
