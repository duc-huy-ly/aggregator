package commands

import (
	"context"
	"fmt"

	"github.com/duc-huy-ly/aggregator/internal/config"
	"github.com/duc-huy-ly/aggregator/internal/database"
)

type State struct {
	MyConfig *config.Config
	Db       *database.Queries
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
		fmt.Printf("Unknown action : %s\n", cmd.Name)
		return fmt.Errorf("Unknown action : %s\n", cmd.Name)
	}

	return fn(s, cmd)
}

func (commands *Commands) Register(name string, f func(*State, Command) error) {
	commands.Map[name] = f
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("commands argument is empty\n")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Something went wrong getting the user")
	}

	fmt.Println("User has been set")
	return s.MyConfig.SetUser(cmd.Args[0])
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("No argument given")
	}
	fmt.Printf("Register function called for %v\n", cmd.Args[0])
	newUser, err := s.MyConfig.RegisterUser(cmd.Args[0], s.Db)
	if err != nil {
		return fmt.Errorf("Something went wrong in the user creation process")
	}
	s.MyConfig.Current_user_name = newUser.Name
	return s.MyConfig.SetUser(newUser.Name)
}

func HandlerReset(s *State, cmd Command) error {
	return s.MyConfig.ResetDatabase(s.Db)
}

func HandlerGetUsers(s *State, cmd Command) error {
	listOfUsers, err := s.MyConfig.GetUsers(s.Db)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	for _, user := range listOfUsers {
		fmt.Printf("* %v", user.Name)
		if user.Name == s.MyConfig.Current_user_name {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")
	}
	return nil
}

func HandlerAggregateDefault(s *State, cmd Command) error {
	return s.MyConfig.Aggregate("https://www.wagslane.dev/index.xml")
}

func HandlerDisplayFeeds(s *State, cmd Command) error {
	return s.MyConfig.DisplayFeeds(s.Db)
}
