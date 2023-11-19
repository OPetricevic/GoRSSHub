package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OPetricevic/GoRSSHub/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv" //Paket - omogućava mi da pozovem Port iz .env

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the enviorment")
	} else {
		fmt.Println("Port:", portString)
	}

	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("DB_URL is not found in the enviorment")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't Connect to database", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	// GET
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleError)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	// POST
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	server.ListenAndServe()
}
