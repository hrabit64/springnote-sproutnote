package ui

import (
	"github.com/manifoldco/promptui"
)

func Password() (string, error) {

	prompt := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}

	return prompt.Run()

}
