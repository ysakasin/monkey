package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Println("The Monkey programming language")
	repl.Start(os.Stdin, os.Stdout)
}
