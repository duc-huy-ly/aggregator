package commands

import (
	"context"
	"fmt"

	"github.com/duc-huy-ly/Gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error{
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Unfollow requires 1 arg : the feed's RL")
	}

	url := cmd.Args[0]
	feed, err := s.Db.GetFeedFromUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("Error : feed not found. %v\n", err)
	}

	params := database.DeleteFeedParams {
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.Db.DeleteFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error deleting feed : %v\n", err)
	}
	return nil

}