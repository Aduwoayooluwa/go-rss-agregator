package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aduwoayooluwa/go-rss-scraper/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	router := chi.NewRouter()

	// ?  mongo db connection
	ctx := context.TODO()

	config := db.MongoDBConfig{
		URI: os.Getenv("MONGODB_URI"),
	}

	if err := db.ConnectMongoDB(ctx, config); err != nil {
		log.Fatalf("Error initializing mongo DB client: %v", err)
	}

	defer db.DisconnectMongoDB(ctx)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//  creating routers
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Get("/get-user-data", handleGetUserData)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}

	fmt.Println("PORT", portString)
}
