package main

import (
	"github.com/google/uuid"
	"github.com/olavowilke/rss-api/internal/database"
	"time"
)

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	ApiKey    string    `json:"apiKey"`
}

func mapDatabaseUserToCreateUserResponse(dbUser database.User) CreateUserResponse {
	return CreateUserResponse{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		CreatedAt: dbUser.CreatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}
