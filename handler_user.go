package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/olavowilke/rss-api/internal/database"
	"log"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  params.Username,
		Email:     params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	respondWithJSON(w, 201, mapDatabaseUserToCreateUserResponse(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, mapDatabaseUserToCreateUserResponse(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		log.Println("Error getting posts for user:", err)
	}

	respondWithJSON(w, 200, mapDatabasePostsToGetPostsResponse(posts))
}
