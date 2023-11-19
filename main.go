package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv" //Paket - omogućava mi da pozovem Port iz .env
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the enviorment")
	} else {
		fmt.Println("Port:", portString)
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
