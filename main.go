package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

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
