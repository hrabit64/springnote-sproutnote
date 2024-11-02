package core

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/cmd"
	"strconv"
)

func processFileExecutor(args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
		return
	}

	switch args[1] {
	case "add":
		cmd.RunFileAdd()
		return
	case "show":
		progressRunFileShow(args)
		return
	case "backup":
		progressFileBackup(args)
		return
	case "remove":
		progressFileRemove(args)
		return
	}
}

func progressRunFileShow(args []string) {
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
	cmd.RunFileShow(page - 1)
	return
}

func progressFileBackup(args []string) {
	if len(args) < 3 {
		cmd.RunFileBackupAll()
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
		cmd.RunFileBackupSingle(target)
		return
	}
}

func progressFileRemove(args []string) {
	if len(args) < 3 {
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid id.")
		return
	}

	cmd.RunFileRm(id)
}
