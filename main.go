package main

import (
	"Interpreter-made-in-Go-language/repl"
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel frree to type in commands\n")
	repl.Start()
}
