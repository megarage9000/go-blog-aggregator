package main

import(
	"fmt"
	"context"
	"github.com/google/uuid"
	"time"
	"github.com/megarage9000/go-blog-aggregator/internal/database"	
)

func addFeed(name string, url string, user database.User, db *database.Queries) error {

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

	follow(db, url, user)

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

func follow(db *database.Queries, url string, user database.User) error {

	feed, err := db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	result, err := db.CreateFeedFollow(context.Background(), feedFollow)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", result)

	return nil
}

func following(db *database.Queries, user database.User) error {

	feedFollows, err := db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("%s\n", feedFollow.FeedName)
	}

	return nil
}

func unfollow(db *database.Queries, user database.User, url string) error {
	feed, err := db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return err
	}

	unfollowParams := database.UnfollowFeedParams {
		UserID: user.ID,
		FeedID: feed.ID,
	}

	return db.UnfollowFeed(context.Background(), unfollowParams)
}