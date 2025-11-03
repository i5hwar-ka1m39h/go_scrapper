package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/i5hwar-ka1m39h/go_scrapper/internal/auth"
	"github.com/i5hwar-ka1m39h/go_scrapper/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) auth_middleware(handle authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			log.Println("error getting api key", err)
			errorResponse(w, 403, fmt.Sprintf("you messed up %v", err))
			return
		}

		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			log.Println("error getting user from db", err)
			errorResponse(w, 500, fmt.Sprintln("error getting user", err))
			return
		}

		handle(w, r, user)
	}
}
