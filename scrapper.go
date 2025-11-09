package main

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/i5hwar-ka1m39h/go_scrapper/internal/database"
)



func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration){
	log.Printf("Scrapping on %v goroutine and %s time between request\n", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C{
		feeds,err := db.GetNotGetFetchedFeeds(context.Background(), int32(concurrency))
		if err != nil{
			log.Println("error occured while fetching feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds{
			wg.Add(1)
			go scrapeshit(wg, db, feed)
		}


		wg.Wait()




	}
}

func scrapeshit(wg *sync.WaitGroup, db *database.Queries, feed database.Feed){
	defer wg.Done()
	_, err:= db.MarkedFeedAsFetched(context.Background(), feed.ID)
	if err != nil{
		log.Println("error marking as fetched!", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil{
		log.Println("error getting the rss feed", err)
		return
	}



	for _, item := range rssFeed.Channel.Item{
		description := sql.NullString{}
		if item.Description != ""{
			description.String = item.Description
			description.Valid = true
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		
		if err != nil{
			log.Println("error parsing the time ", err)
			continue
		}

		post, err := db.CreateFeedPost(context.Background(), database.CreateFeedPostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID: feed.ID,
			Title: item.Title,
			Description: description,
			Url: item.Link,
			PublishedAt: t,
		})

		if err != nil{
			if strings.Contains(err.Error(), "duplicate key"){
				continue
			}
			log.Println("error creating the post", err)
			continue
		}
		log.Println("post created with id: ",post.ID)
	}

	log.Printf("feed %s collected %v items found.", rssFeed.Channel.Title, len(rssFeed.Channel.Item))
}