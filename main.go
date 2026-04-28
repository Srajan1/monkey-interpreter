package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Srajan1/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hemloo %s👋, This is the Monkey programming language, no monkey business around here!!\n", user.Username)
	fmt.Println("Feel free to start typing commands")

	repl.Start(os.Stdin, os.Stdout)
}
