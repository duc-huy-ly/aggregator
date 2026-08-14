package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/duc-huy-ly/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, currentUser database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Follow requires one argument")
	}
	url := cmd.Args[0]

	feed, err := s.Db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("GetFeedFromUrl error : %v\n", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("CreateFeedFollow error : %v\n", err)
	}
	fmt.Printf("Feed name: %v\nUser: %v\n", feedFollow.FeedName, feedFollow.UserName)
	return nil

}
