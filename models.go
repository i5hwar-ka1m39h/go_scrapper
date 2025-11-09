package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/i5hwar-ka1m39h/go_scrapper/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uname     string    `json:"user_name"`
	APIkey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fname     string    `json:"feed_name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

type FeedFollows struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId uuid.UUID `json:"user_id"`
	FeedId uuid.UUID `json:"feed_id"`
}

type Post struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
    Title string `json:"title"` 
    Description *string `json:"description"` 
    Published_at time.Time `json:"published_at"`
    Url string `json:"url"`
	Feed_id uuid.UUID `json:"feed_id"`
}

func dbUserToMdUSer(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Uname:     dbUser.Uname,
		APIkey:    dbUser.ApiKey,
	}
}

func dbFeedToMdFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Fname:     dbFeed.Fname,
		Url:       dbFeed.Url,
		UserId:    dbFeed.UserID,
	}
}

func dbFeedsToMdFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dbFeedToMdFeed(dbFeed))
	}
	return feeds
}

func dbFeedFollowsToMdFeedFollows(dbFeedFollows database.FeedsFollow) FeedFollows{
	return FeedFollows{
		ID: dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserId: dbFeedFollows.UserID,
		FeedId: dbFeedFollows.FeedID,
	}
}

func dbMultFollowsToMultFeedFollows(dbFeedFollowArr []database.FeedsFollow) []FeedFollows{
	feedFollows := []FeedFollows{}
	for _, dbFeedFollows := range dbFeedFollowArr{
		feedFollows = append(feedFollows, dbFeedFollowsToMdFeedFollows(dbFeedFollows))
	}
	return feedFollows
}

func dbPostToMdPost(dbPost database.Post) Post{
	var description *string
	if dbPost.Description.Valid{
		description = &dbPost.Description.String
	}
	return  Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title: dbPost.Title,
		Description: description,
		Url: dbPost.Url,
		Feed_id: dbPost.FeedID, 
	}
}

func dbPostsToMdPosts(dbPostArr []database.Post) []Post{
	posts := []Post{}
	for _, post:= range dbPostArr{
		posts = append(posts, dbPostToMdPost(post))
	}

	return posts
}