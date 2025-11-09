package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/i5hwar-ka1m39h/go_scrapper/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("building a go scrapper")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not found")
	}
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("db url not found")
	}
	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("error occured while connection to database", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}


	go startScrapping(db, 10, time.Minute)

	r := chi.NewRouter()

	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"http://*", "https://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1router := chi.NewRouter()

	v1router.Get("/ready", handleSome)
	v1router.Get("/err", handleError)
	v1router.Post("/users", apiCfg.handleCreateUser)
	v1router.Get("/users", apiCfg.auth_middleware(apiCfg.handleGetUser))
	v1router.Post("/feed", apiCfg.auth_middleware(apiCfg.handleCreateFeed))
	v1router.Get("/feeds", apiCfg.handleGetFeeds)
	v1router.Post("/followfeed", apiCfg.auth_middleware(apiCfg.handleCreateFeedFollows))
	v1router.Get("/followfeed", apiCfg.auth_middleware(apiCfg.handleGetFeedFollows))
	v1router.Delete("/followfeed/{feedFollowId}", apiCfg.auth_middleware(apiCfg.handleDeleteFeedFollows))
	v1router.Get("/posts", apiCfg.auth_middleware(apiCfg.handleGetPostForUser))

	r.Mount("/v1", v1router)

	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	log.Printf("server running of %v", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("server failed", err)
	}
}
