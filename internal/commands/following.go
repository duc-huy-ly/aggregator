package commands

import (
	"context"
	"fmt"

	"github.com/duc-huy-ly/aggregator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	userName := s.MyConfig.Current_user_name

	feedsOfuser, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error fetching the feeds of user %v : %v\n", userName, err)
	}
	for i, v := range feedsOfuser {
		fmt.Printf("Feed #%v : '%v'\n", i+1, v.FeedName)
	}

	return nil
}
