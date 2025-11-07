package main

import (
	"time"

	"github.com/i5hwar-ka1m39h/go_scrapper/internal/database"
)



func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration)