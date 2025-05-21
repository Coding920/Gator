package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Coding920/Gator/internal/config"
	"github.com/Coding920/Gator/internal/database"
	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}
	cmds := commands{commands: make(map[string]func(*state, command) error)}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetUsersDb)
	cmds.register("users", listUsers)

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
