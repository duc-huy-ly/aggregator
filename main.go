package main
import _ "github.com/lib/pq"

import (
	"fmt"
	"os"
	"github.com/duc-huy-ly/aggregator/internal/commands"
	"github.com/duc-huy-ly/aggregator/internal/config"
)

func main() {
	fmt.Printf("Welcome to blog aggretator\n")
	localConfig := config.Read()
	localState := commands.State{
		MyConfig: &localConfig,
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
	err := allCommands.Run(&localState, myCommand)
	if err != nil {
		os.Exit(1)
	}
}
