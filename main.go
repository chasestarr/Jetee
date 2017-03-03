package main

import (
	"fmt"
	"github.com/chasestarr/Jetee/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Jetée programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
