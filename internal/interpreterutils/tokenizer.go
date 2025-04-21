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

func parseTokens(tokens []token) ([]int, error) {
	var memory [30000]int
	pointer := 0
	currentSquareBracket := "None"

	for _, item := range tokens {
		if item.symbol == "+" {
			memory[pointer]++
		}
		if item.symbol == "-" {
			memory[pointer]--
		}
		if item.symbol == ">" {
			if pointer < 30000 {
				pointer++
			} else {
				//todo throw err
			}
		}
		if item.symbol == "<" {
			if pointer > 0 {
				pointer--
			} else {
				//todo throw err
			}
		}
		if item.symbol == "[" {
			currentSquareBracket = item.symbol
			if memory[pointer] == 0 {

			}
		}
	}

}
