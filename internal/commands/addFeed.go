package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/duc-huy-ly/aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("AddFeed(name string, url string)")
	}
	return AddFeed(cmd.Args[1], cmd.Args[0], s.MyConfig.Current_user_name, s.Db)
}

func AddFeed(url string, name string, currentUsername string, db *database.Queries) error {
	if url == "" || name == "" {
		return fmt.Errorf("config.AddFeed(): Arguments must not be null")
	}
	currentUser, err := db.GetUser(context.Background(), currentUsername )
	if err != nil {
		return err
	}
	newFeedParams := database.CreateFeedParams{
		ID:uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: currentUser.ID,
	}
	feed, err := db.CreateFeed(context.Background(), newFeedParams)	
	if err != nil {
		return err
	}

	fmt.Printf("ID : %v\nName : %v\nCreated at : %v\nUpdated at : %v\nURL : %v\nUserID : %v\n", feed.ID, feed.Name, feed.CreatedAt, feed.UpdatedAt, feed.Url, feed.UserID)

	// create feed follow record
	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}
	_ , err = db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Create feed follow error inside AddFeed : %v\n", err)
	}
	return nil
}