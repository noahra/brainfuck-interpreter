package interpreterutils

import (
	"bufio"
	"fmt"
	"os"
)

func RunRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nBF-REPL> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				break
			}
			tokens := interpretInput(input)
			fmt.Println()
			output, _ := executeBrainfuck(tokens)
			for _, char := range output {
				fmt.Printf("%c", char)
			}
		}
	}
}
