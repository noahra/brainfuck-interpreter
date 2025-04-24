package interpreterutils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func executeBrainfuck(tokens []token) (string, error) {
	var memory [30000]int
	var output []rune
	pointer := 0

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].symbol {
		case "+":
			memory[pointer]++
		case "-":
			memory[pointer]--
		case ">":
			if pointer >= 30000 {
				return "", errors.New("pointer out of bounds: exceeded maximum memory address")
			}
			pointer++
		case "<":
			if pointer <= 0 {
				return "", errors.New("pointer out of bounds: negative memory address")
			}
			pointer--
		case "[":
			if memory[pointer] == 0 {
				i = handleLeftBracket(tokens, i)
			}
		case "]":
			if memory[pointer] != 0 {
				i = handleRightBracket(tokens, i)
			}
		case ".":
			output = append(output, rune(memory[pointer]))

		case ",":
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				memory[i] = int(scanner.Bytes()[0])
			}
		}

	}

	return strings.ReplaceAll(string(output), "\n", ""), nil
}
func handleRightBracket(tokens []token, outerIndex int) int {
	var squareBrackets Stack
	squareBrackets.Push("[")
	for j := outerIndex - 1; j > 0; j-- {
		if tokens[j].symbol == "[" {
			squareBrackets.Pop()
		}
		if tokens[j].symbol == "]" {
			squareBrackets.Push("]")
		}
		if squareBrackets.IsEmpty() {
			return j - 1
		}
	}
	return outerIndex
}
func handleLeftBracket(tokens []token, outerIndex int) int {
	var squareBrackets Stack
	squareBrackets.Push("[")
	for j := outerIndex + 1; j < len(tokens); j++ {
		if tokens[j].symbol == "]" {
			squareBrackets.Pop()
		}
		if tokens[j].symbol == "[" {
			squareBrackets.Push("[")
		}
		if squareBrackets.IsEmpty() {
			return j
		}
	}
	return outerIndex
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
