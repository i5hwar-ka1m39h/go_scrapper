package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
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

func (cfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		log.Println("error getting feeds follows", err)
		errorResponse(w, 404, fmt.Sprintln("can't get the feeds follows", err))
		return
	}

	jsonResponseWriter(w, 200, dbMultFollowsToMultFeedFollows(feedsFollows))
}

type DeleteResp struct{
	Message string `json:"message"`
}
func (cfg *apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollowID := chi.URLParam(r, "feedFollowId")
	ffID, err := uuid.Parse(feedFollowID)
	

	if err != nil{
		log.Println("error parsing the feed id", err)
		errorResponse(w, 500, fmt.Sprintln("error getting id",err))
		return
	}


	err = cfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID: ffID,
		UserID: user.ID,
	})

	if err != nil{
		log.Println("error deleteing feed follow", err)
		errorResponse(w, 500, fmt.Sprintln("error deleting the feed folow", err))
		return
	}

	resp := DeleteResp{
		Message: "unfollowd the feed with given id.",
	}
	jsonResponseWriter(w, 200, resp)

}