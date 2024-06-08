package main

import (
	"log"
	"net/http"

	"github.com/aduwoayooluwa/go-rss-scraper/db"
	"github.com/aduwoayooluwa/go-rss-scraper/models"
)

func handleGetUserData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := db.GetAllUsers(ctx)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve users: "+err.Error())
	}

	respondWithJSON(w, http.StatusOK, users)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       30,
	}

	if err := db.CreateUser(ctx, user); err != nil {
		log.Fatalf("failed to create user: %v", err)
		respondWithError(w, 500, "Something went Wrong")
	}

	respondWithJSON(w, http.StatusOK, "Saved to the DB successfully")
}
