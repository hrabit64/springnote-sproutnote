package core

import (
	"github.com/c-bata/go-prompt"
)

func RunPrompt() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionHistory([]string{}),
		prompt.OptionTitle("SproutNote Cil"),
	)
	p.Run()
}
