package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/duc-huy-ly/Gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) != 0 {
		mylimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Parsing value err")
		}
		limit = mylimit
	}

	posts, err := s.Db.GetPostsForUserInt(context.Background(), limit)
	if err != nil {
		return fmt.Errorf("HandlerBrowse() : Cant' get posts for user ")
	}
	for _, post := range posts {
		fmt.Printf("%s\n%s\n\n", post.Title, post.Url)
	}

	return nil

}
