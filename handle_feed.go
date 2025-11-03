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

func (cfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("error parsing the json ", err)
		errorResponse(w, 400, fmt.Sprintln("error parsing json", err))
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
		Fname:     params.Name,
	})
	if err != nil {
		log.Println("error creating a feed", err)
		errorResponse(w, 500, fmt.Sprintf("we fucked up %v", err))
		return
	}

	jsonResponseWriter(w, 201, dbFeedToMdFeed(feed))
}

func (cfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Println("error getting feeds", err)
		errorResponse(w, 404, fmt.Sprintln("can't get the feeds", err))
		return
	}

	jsonResponseWriter(w, 200, dbFeedsToMdFeeds(feeds))
}
