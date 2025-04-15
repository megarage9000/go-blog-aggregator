package main

import(
	"context"
	"net/http"
	"io"
	"encoding/xml"
	"html"
	"fmt"
	"github.com/megarage9000/go-blog-aggregator/internal/database"
	"time"
)

// Helper function for RSSFeed and RSSItem feed
func (rssItem *RSSItem) PrintFeedItem() {
	fmt.Println(" ---- RSS Item ---- ")
	fmt.Println(rssItem.Title)
	fmt.Println(rssItem.Link)
	fmt.Println(rssItem.Description)
	fmt.Println(rssItem.PubDate)
}

func (rssFeed *RSSFeed) PrintFeed() {
	fmt.Println("---- RSS Feed ---- ")
	fmt.Println(rssFeed.Channel.Title)
	fmt.Println(rssFeed.Channel.Link)
	fmt.Println(rssFeed.Channel.Description)

	for _, rssItem := range rssFeed.Channel.Item {
		rssItem.PrintFeedItem()
	}
}

func scrapeFeed(ctx context.Context, db *database.Queries) error {
	feed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("GetNextFeedFetch: %w", err)
	}

	markFeedFetchedArgs := database.MarkFeedFetchedParams {
		ID: feed.ID,
		UpdatedAt: time.Now(),
	}

	feedFetchedErr := db.MarkFeedFetched(ctx, markFeedFetchedArgs)
	if feedFetchedErr != nil {
		return fmt.Errorf("MarkFeedFetched: %w", feedFetchedErr)
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("fetchFeed: %w", feedFetchedErr)
	}

	savePostErr := savePost(ctx, db, rssFeed, feed.ID)
	if savePostErr != nil {
		return savePostErr
	}
	return nil
}


func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	// Doing a Get Request
	req, err:= http.NewRequestWithContext(ctx, "GET", feedUrl, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	// Performing the Get Request
	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		return nil, err
	}

	// Reading the data
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)

	var results RSSFeed
	if xml.Unmarshal(bytes, &results) != nil {
		return nil, err
	}

	// Clean the string input
	RemoveEntities(&results.Channel.Title)
	RemoveEntities(&results.Channel.Description)

	for i := 0; i < len(results.Channel.Item); i++ {
		RemoveEntities(&results.Channel.Item[i].Title)
		RemoveEntities(&results.Channel.Item[i].Description)
	}

	return &results, nil
}

func RemoveEntities(str *string) {
	*str = html.UnescapeString(*str)
}



