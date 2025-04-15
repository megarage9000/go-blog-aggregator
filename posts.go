package main

import(
	"context"
	"github.com/google/uuid"
	"fmt"
	"github.com/megarage9000/go-blog-aggregator/internal/database"
	"time"
	"database/sql"
)


func savePost(ctx context.Context, db * database.Queries, rssFeed * RSSFeed, feedID uuid.UUID) error {

	for _, rssItem := range rssFeed.Channel.Item {

		var rssDescription sql.NullString
		var rssPubDate sql.NullTime

		rssDescription.Scan(rssItem.Description)
		rssPubDate.Scan(rssItem.PubDate)
	
		postArgs := database.CreatePostParams {
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: rssItem.Title,
			Url: rssItem.Link,
			Description: rssDescription,
			PublishedAt: rssPubDate,
			FeedID: feedID,
		}
		
		fmt.Printf("--- Creating post for %s ---\n", postArgs.Title)
		_, err := db.CreatePost(ctx, postArgs)



		if err != nil {
			if err != sql.ErrNoRows {
				return fmt.Errorf("--- error in creating post for %s: Error code %s = %s --- \n", postArgs.Title, err, err)
			} else {
				fmt.Printf("--- Post %s is duplicate --- \n", postArgs.Title)
			}
		}
	}

	return nil
}

func browsePosts(ctx context.Context, user database.User, db * database.Queries, numItems int32) error {

	userPostsArgs := database.GetPostsForUserParams {
		UserID: user.ID,
		Limit: numItems,
	}

	userPosts, err := db.GetPostsForUser(ctx, userPostsArgs)
	if err != nil {
		return err
	}

	for _, userPost := range userPosts {
		var description string
		var pubDate time.Time
		if userPost.Description.Valid {
			description = userPost.Description.String
		}
		if userPost.PublishedAt.Valid {
			pubDate = userPost.PublishedAt.Time
		}
		fmt.Printf("--- Post --- \nTitle: %s\nDescription: %s\nPublished At: %v\n", userPost.Title, description, pubDate)
	}

	return nil
}