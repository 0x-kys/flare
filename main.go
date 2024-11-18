package main

import (
	"flare/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("[", user.Username, "] @ Flare")
	fmt.Println("Start typing to get started")
	fmt.Println("Input `.bye` to exit")

	repl.Start(os.Stdin, os.Stdout)
}
