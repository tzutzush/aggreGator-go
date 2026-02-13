package main

import (
	"aggreGATOR/internal/config"
	"aggreGATOR/internal/database"
	"database/sql"
	"fmt"
	"os"
)
import _ "github.com/lib/pq"

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("reading config unsuccessful")
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	applicationState := &state{config: &cfg, db: dbQueries}
	cmds := commandRegistry{commands: make(map[string]func(*state, command) error)}
	cmds.register("login", loginHandler)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)

	if len(os.Args) < 2 {
		fmt.Println("no argument was provided")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(applicationState, command{name: cmdName, arguments: cmdArgs})
	if err != nil {
		fmt.Printf("err was: %v\n", err)
		os.Exit(1)
	}
}
