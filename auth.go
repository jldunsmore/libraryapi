package main

import (
	c "context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userId = r.URL.Query().Get("userId")
		user, ok := userDatabase[userId]
		if !ok || userId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token := r.Header.Get("Authorization")
		if !isValidToken(user, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := c.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	}
}

func isValidToken(user User, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == user.Token
	}
	return false
}
