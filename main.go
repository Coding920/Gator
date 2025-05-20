package main

import (
	"fmt"
	"github.com/Coding920/Gator/internal/config"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide command")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := state{cfg: &cfg}
	cmds := commands{commands: make(map[string]func(*state, command) error)}
	cmds.register("login", loginHandler)

	inputCmd := command{
		name: strings.ToLower(os.Args[1]),
		args: os.Args[2:],
	}

	err = cmds.run(&s, inputCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
