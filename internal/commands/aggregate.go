package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/duc-huy-ly/Gator/internal/database"
	"github.com/duc-huy-ly/Gator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func isDuplicatePostError(err error) bool {
	pqErr, ok := err.(*pq.Error)
	return ok && string(pqErr.Code) == "23505"
}

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
	for _, v := range feed.Channel.Item {
		// fmt.Printf("#%v: %v\n",i+1, v.Title)
		pubdate, err := time.Parse(time.RFC1123Z, v.PubDate)
		if err != nil {
			return fmt.Errorf("Error parsing the time of %v. %v", v, err)
		}

		params := database.CreatepostParams{
			ID:          uuid.New(),
			Title:       v.Title,
			Url:         v.Link,
			Description: v.Description,
			PublishedAt: pubdate,
			FeedID:      nextFeed.ID,
		}
		_, err = s.Db.Createpost(context.Background(), params)
		if err != nil {
			if isDuplicatePostError(err) {
				continue
			}
			return fmt.Errorf("create post: %w", err)
		}

	}

	return nil

}

func HandlerAggregate(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Aggretate needs time_between_reqs")
	}
	frequency, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error parsing time duration")
	}

	fmt.Printf("Collecting feeds every %v\n", frequency)

	ticker := time.NewTicker(frequency)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return err
		}
		fmt.Printf("\n")
	}

}
