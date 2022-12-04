package prompt

import (
	"errors"
	"github.com/manifoldco/promptui"
)

func Validate(input string) error {
	if input == "" {
		return errors.New("please enter a valid domain")
	}
	return nil
}

func CreatePrompt(label string, v func(string) error) *promptui.Prompt {

	prompt := promptui.Prompt{
		Label:    label,
		Validate: v,
	}
	return &prompt
}

func CreateSelectPrompt(label string, vals []string) *promptui.Select {
	prompt := promptui.Select{
		Label: label,
		Items: vals,
	}
	return &prompt
}
