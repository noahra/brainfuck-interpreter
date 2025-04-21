package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
