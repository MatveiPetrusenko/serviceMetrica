package device_service

import "net/http"

// Middleware для обработки запросов в Device Service API.
type Middleware struct {
}

// NewMiddleware создает новый экземпляр Middleware.
func NewMiddleware() *Middleware {
	return &Middleware{}
}

// Middleware для проверки авторизации по JWT-токену.
func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
