package core

import (
	"github.com/hrabit64/sproutnote/cil/cmd"
	"strings"
)

func executor(input string) {
	input = strings.TrimSpace(input)

	if input == "" {
		cmd.RunHelp()
	}

	args := strings.Split(input, " ")

	switch args[0] {

	case "help":
		cmd.RunHelp()

	case "exit":
		cmd.RunExit()
		return

	case "db":
		processDbExecutor(args)
		return

	case "file":
		processFileExecutor(args)
		return

	case "history":
		processHistoryExecutor(args)
		return

	case "config":
		processConfigExecutor(args)
		return

	default:
		cmd.RunHelp()
		return
	}
}
