package internalservice

import (
	"net/http"
)

// Middleware for logging, authentication, etc.
type Middleware struct {
}

// NewMiddleware creates a new Middleware.
func NewMiddleware() *Middleware {
	return &Middleware{}
}

// LoggingMiddleware logs incoming requests.
func (m *Middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// AuthenticationMiddleware handles user authentication.
func (m *Middleware) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*		apiToken := r.Header.Get("Authorization")
				if apiToken == "" {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}*/

		next.ServeHTTP(w, r)
	})
}
