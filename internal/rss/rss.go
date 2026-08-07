package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// Fetch a given URL, assuming nothing goes wrong, return a filled out RSSFeed struct
func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	feed := RSSFeed{}
	if feedURL == "" {
		return nil, fmt.Errorf("URL must not be empty")
	}
	getRequestFeedUrl, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error while making the GET request : %v\n", err)
	}
	getRequestFeedUrl.Header.Set("User-Agent", "gator")
	myClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	myResponse, err := myClient.Do(getRequestFeedUrl)	
	if err != nil {
		return nil, fmt.Errorf("Error fetching the response : %v\n", err)
	}
	defer myResponse.Body.Close()

	body, err := io.ReadAll(myResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("Error decoding the body : %v\n", err)
	}

	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling body : %v\n", err)
	}

	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}


	return &feed, nil

}