package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/duc-huy-ly/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, currentUser database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("AddFeed(name string, url string)")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	db := s.Db

	newFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}

	feed, err := db.CreateFeed(context.Background(), newFeedParams)
	if err != nil {
		return err
	}

	fmt.Printf("ID : %v\nName : %v\nCreated at : %v\nUpdated at : %v\nURL : %v\nUserID : %v\n", feed.ID, feed.Name, feed.CreatedAt, feed.UpdatedAt, feed.Url, feed.UserID)

	// create feed follow record
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	_, err = db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Create feed follow error inside AddFeed : %v\n", err)
	}
	return nil
}
