package main

import (
	"fmt"
	"os"
	"os/user"

	"lucasrego.tech/monkey-lang/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("[EN] Hello %s! This is the Monkey programming language! Feel free to type commands\n", user.Username)
	fmt.Printf("[PT] Olá %s! Está é a linguagem de programação Monkey! Sinta-se livre para digitar comandos\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
