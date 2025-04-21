package main

import "fmt"

func main() {

	for {
		fmt.Print("BF-REPL> ")
		var input string
		_, err := fmt.Scan(&input)
		if input == "exit" {
			break
		}
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Println(input)
	}

}
