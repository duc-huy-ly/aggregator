package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/duc-huy-ly/aggregator/internal/config"
	"github.com/duc-huy-ly/aggregator/internal/database"
	"github.com/google/uuid"
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

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("AddFeed(name string, url string)")
	}
	return s.MyConfig.AddFeed(cmd.Args[1], cmd.Args[0], s.Db)
}

func HandlerDisplayFeeds(s *State, cmd Command) error {
	return s.MyConfig.DisplayFeeds(s.Db)
}

func HandlerFollow(s *State, cmd Command)error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Follow requires one argument")
	}
	url := cmd.Args[0]
	feed, err := s.Db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("GetFeedFromUrl error : %v\n", err)
	}

	currentUser, err := s.Db.GetUser(context.Background(), s.MyConfig.Current_user_name)
	if err != nil {
		return fmt.Errorf("Error fetching user ID : %v\n", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}
	feedFollow, err:= s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("CreateFeedFollow error : %v\n", err)
	}
	fmt.Printf("Feed name: %v\nUser: %v\n", feedFollow.FeedName, feedFollow.UserName)
	return  nil

}

func HandlerFollowing(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Error : following command requires 1 argument, the name of the given user")
	}
	userName := cmd.Args[0]
	selectedUser, err:= s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error fetching the user : %v\n")
	}
	feedsOfuser, err := s.Db.GetFeedFollowsForUser(context.Background(), selectedUser.ID)
	if err != nil {
		return fmt.Errorf("Error fetching the feeds of user %v : %v\n", userName, err)
	}
	for i, v := range feedsOfuser {
		fmt.Printf("Feed #%v : '%v'\n", i + 1, v.FeedName)
	}
	
	return nil
}