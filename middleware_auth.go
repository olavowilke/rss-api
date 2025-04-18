package main

import (
	"github.com/olavowilke/rss-api/internal/auth"
	"github.com/olavowilke/rss-api/internal/database"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, authedUser database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, "user not found")
			return
		}

		handler(w, r, user)
	}
}
