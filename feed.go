package main

import(
	"fmt"
	"context"
	"github.com/google/uuid"
	"time"
	"github.com/megarage9000/go-blog-aggregator/internal/database"	
)

func addFeed(name string, url string, username string, db *database.Queries) error {

	user, err := db.GetUser(context.Background(), username)

	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	}

	feed, err := db.CreateFeed(context.Background(), feedParams)

	if err != nil {
		return err
	}

	fmt.Printf("%+v", feed)
	return nil
}

func listFeed(db *database.Queries) error {
	feeds, err := db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		name := feed.Name
		url := feed.Url
		userName, err := db.GetUserName(context.Background(), feed.UserID)
		if err == nil {
			fmt.Printf(" ==== FEED ==== \n%s\n%s\n%s\n", name, url, userName)
		}
	}

	return nil
}