package interpreterutils

import "strings"

func interpretInput(input string) []token {
	var tokens []token
	validCommands := "+-><[].,"
	for _, ch := range input {
		cmd := string(ch)
		if strings.Contains(validCommands, cmd) {
			tokens = append(tokens, token{cmd})
		}
	}
	return tokens
}

type token struct {
	symbol string
}
