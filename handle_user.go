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

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("error parsing the json ", err)
		errorResponse(w, 400, fmt.Sprintln("error parsing json", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Uname:     params.Name,
	})
	if err != nil {
		log.Println("error creating a user", err)
		errorResponse(w, 500, fmt.Sprintf("we fucked up %v", err))
		return
	}

	jsonResponseWriter(w, 201, dbUserToMdUSer(user))
}

func (apiConfig *apiConfig) handleGetUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	jsonResponseWriter(w, 200, dbUserToMdUSer(user))
}
