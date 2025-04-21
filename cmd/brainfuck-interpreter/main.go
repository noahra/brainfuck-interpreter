package main

import (
	"github.com/noahra/brainfuck-interpreter/internal/interpreterutils"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		interpreterutils.HandleScript()
	} else {
		interpreterutils.RunRepl()
	}
}
