package interpreterutils

import (
	"fmt"
	"os"
)

func HandleScript() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error occurred: ", err)
	}
	tokens := interpretInput(string(file))
	fmt.Println()
	executeBrainfuck(tokens)
}
