package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aduwoayooluwa/go-rss-scraper/db"
	"github.com/aduwoayooluwa/go-rss-scraper/models"
	"github.com/go-chi/chi"
)

func handleGetUserData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//users, err := db.GetAllUsers(ctx)

	userId := chi.URLParam(r, "userId")

	if userId == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// retreiving user info
	user, err := db.GetUserById(ctx, userId)

	if err != nil {
		if strings.Contains(err.Error(), "invalid user ID format") {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err.Error() == fmt.Sprintf("no user found with ID %s", userId) {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve users: "+err.Error())
		return
	}
	// fmt.Printf("userId %v \n", userId)
	// fmt.Printf("user info %v", user)
	respondWithJSON(w, http.StatusOK, user)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	newUser := models.User{}

	decoder := json.NewDecoder(r.Body)

	decoderError := decoder.Decode(&newUser)

	if decoderError != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json"))
		return
	}

	//  closing the request body
	defer r.Body.Close()

	if err := db.CreateUser(ctx, newUser); err != nil {
		log.Fatalf("failed to create user: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error creating user: "+err.Error())
	}

	respondWithJSON(w, http.StatusCreated, newUser)
}
