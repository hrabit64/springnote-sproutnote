package ui

import "github.com/manifoldco/promptui"

func InputBox(label string) (string, error) {

	prompt := promptui.Prompt{
		Label: label,
	}

	return prompt.Run()

}
