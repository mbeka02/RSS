package main

import (
	"fmt"
	"net/http"

	"github.com/mbeka02/RSS/internal/auth"

	"github.com/mbeka02/RSS/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

// method
func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	//return anon func with handler func signature
	return func(w http.ResponseWriter, r *http.Request) {
		authKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			errorResponse(w, 403, fmt.Sprintf("Invalid auth credentials: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), authKey)

		handler(w, r, user)
		if err != nil {
			errorResponse(w, 400, fmt.Sprintf("Unable to find user credentials: %v", err))
			return
		}
	}
}
