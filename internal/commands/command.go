package commands

import (
	"fmt"

	"github.com/duc-huy-ly/aggregator/internal/config"
	"github.com/duc-huy-ly/aggregator/internal/database"
)

type State struct {
	MyConfig *config.Config
	Db *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Map map[string]func(*State, Command) error
}

func (commands *Commands) Run(s *State, cmd Command) error {
	if s == nil {
		return fmt.Errorf("State is nil. \n")
	}
	fn, ok := commands.Map[cmd.Name]
	if !ok {
		return fmt.Errorf("Unknown action : %s\n", cmd.Name)
	}

	return fn(s, cmd)
}

func (commands *Commands) Register (name string, f func(*State, Command) error){
	commands.Map[name] = f
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("commands argument is empty\n")
	}
	fmt.Println("User has been set")
	return s.MyConfig.SetUser(cmd.Args[0])
}


