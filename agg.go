package main

import(
	"context"
	"net/http"
	"io"
)


func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	req, err:= http.NewRequestWithContext(ctx, "GET", feedUrl, nil)

	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)

	return nil, nil
}