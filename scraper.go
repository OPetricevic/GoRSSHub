package main

import (
	"log"
	"time"

	"github.com/OPetricevic/GoRSSHub/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Print("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

}
