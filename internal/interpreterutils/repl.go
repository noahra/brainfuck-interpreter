package interpreterutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunRepl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nBF-REPL> ")
		var lines []string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			if len(strings.TrimSpace(line)) == 0 {
				break
			}
			lines = append(lines, line)
		}
		input := strings.Join(lines, "")
		tokens := interpretInput(input)
		fmt.Println()
		output, _ := executeBrainfuck(tokens)
		fmt.Println(output)

	}
}
