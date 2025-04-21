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

func parseTokens(tokens []token) ([30000]int, error) {
	var memory [30000]int
	pointer := 0

	for i := 0; i < len(tokens); i++ {
		item := tokens[i]
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
			if memory[pointer] == 0 {
				var squareBrackets Stack
				squareBrackets.Push(item.symbol)
				for j := i + 1; j < len(tokens); j++ {
					if tokens[j].symbol == "]" {
						squareBrackets.Pop()
					}
					if tokens[j].symbol == "[" {
						squareBrackets.Push("[")
					}
					if squareBrackets.IsEmpty() {
						i = j
						break
					}

				}

			}
		}
		if item.symbol == "]" {
			if memory[pointer] != 0 {
				var squareBrackets Stack
				squareBrackets.Push(item.symbol)
				for j := i - 1; j > 0; j-- {
					if tokens[j].symbol == "[" {
						squareBrackets.Pop()
					}
					if tokens[j].symbol == "]" {
						squareBrackets.Push("]")
					}
					if squareBrackets.IsEmpty() {
						i = j - 1
						break
					}
				}

			}
		}
	}

	return memory, nil
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
