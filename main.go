package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/duc-huy-ly/aggregator/internal/commands"
	"github.com/duc-huy-ly/aggregator/internal/config"
	"github.com/duc-huy-ly/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("Welcome to blog aggretator\n")
	localConfig := config.Read()

	db, err := sql.Open("postgres", localConfig.DatabaseURL())
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		fmt.Printf("database connection failed: %v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	localState := commands.State{
		MyConfig: &localConfig,
		Db:       dbQueries,
	}

	allCommands := commands.Commands{
		Map: make(map[string]func(*commands.State, commands.Command) error),
	}
	allCommands.Register("login", commands.HandlerLogin)
	allCommands.Register("register", commands.HandlerRegister)
	allCommands.Register("reset", commands.HandlerReset)

	if len(os.Args) < 2 {
		fmt.Printf("Error, less than 2 arguments given\n")
		os.Exit(1)
	}

	args := os.Args[1:]
	myCommand := commands.Command{
		Name: args[0],
		Args: args[1:],
	}
	err = allCommands.Run(&localState, myCommand)
	if err != nil {
		os.Exit(1)
	}
}
