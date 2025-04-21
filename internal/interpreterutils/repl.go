package interpreterutils

import (
	"bufio"
	"fmt"
	"os"
)

func RunRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("BF-REPL> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				break
			}

			fmt.Println(input)
		}
	}
}
