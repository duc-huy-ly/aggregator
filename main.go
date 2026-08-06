package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/duc-huy-ly/aggregator/internal/commands"
	"github.com/duc-huy-ly/aggregator/internal/config"
	"github.com/duc-huy-ly/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("Welcome to blog aggretator\n")
	localConfig := config.Read()

	db, err := sql.Open("postgres", localConfig.Db_url)
	if err != nil {
		os.Exit(1)
	}

	dbQueries := database.New(db)

	localState := commands.State{
		MyConfig: &localConfig,
		Db : dbQueries,
	}

	allCommands := commands.Commands{
		Map: make(map[string]func(*commands.State, commands.Command) error),
	}
	allCommands.Register("login", commands.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Error, less than 2 arguments given\n")
		os.Exit(1)
	}

	args := os.Args[1:]
	myCommand := commands.Command{
		Name: args[0],
		Args: args[1:],
	}
	fmt.Println(args[0])
	fmt.Println(args[1:])
	err = allCommands.Run(&localState, myCommand)
	if err != nil {
		os.Exit(1)
	}
}
