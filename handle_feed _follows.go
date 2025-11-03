package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/i5hwar-ka1m39h/go_scrapper/internal/database"
)

func (cfg *apiConfig) handleCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("error parsing the json ", err)
		errorResponse(w, 400, fmt.Sprintln("error parsing json", err))
		return
	}

	feedfollows, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		
		UserID:    user.ID,
		FeedID: params.FeedId,
	})
	if err != nil {
		log.Println("error creating a feed follow", err)
		errorResponse(w, 500, fmt.Sprintf("we fucked up %v", err))
		return
	}

	jsonResponseWriter(w, 201, dbFeedFollowsToMdFeedFollows(feedfollows))
}

// func (cfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request) {
// 	feeds, err := cfg.DB.GetFeeds(r.Context())
// 	if err != nil {
// 		log.Println("error getting feeds", err)
// 		errorResponse(w, 404, fmt.Sprintln("can't get the feeds", err))
// 		return
// 	}

// 	jsonResponseWriter(w, 200, dbFeedsToMdFeeds(feeds))
// }
