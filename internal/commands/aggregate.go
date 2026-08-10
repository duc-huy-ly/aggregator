package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/duc-huy-ly/aggregator/internal/rss"
)

func scrapeFeeds(s *State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("scrapefeeds err : %v\n", err)
	}
	s.Db.MarkFeedFetched(context.Background(), nextFeed.ID)
	feed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Err Fetchfeed : %v\n", err)
	}
	for i, v := range feed.Channel.Item {
		fmt.Printf("#%v: %v\n",i+1, v.Title)
	}
	return nil

}

func HandlerAggregate(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Aggretate needs time_between_reqs")
	}
	frequency , err:= time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error parsing time duration")
	}

	fmt.Printf("Collecting feeds every %v\n", frequency )

	ticker := time.NewTicker(frequency)

	for ;; <-ticker.C {
		scrapeFeeds(s)
		fmt.Printf("\n")
	}

}
