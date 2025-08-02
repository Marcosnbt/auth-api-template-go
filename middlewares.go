// Middleware criado em maio de 2025
// JWT authentication middleware
package middlewares

import (
	"net/http"
	"strings"
)

// AuthMiddleware simula uma verificação de token JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Aqui seria feita a validação do token real
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" || token != "seu-token-de-teste" {
			http.Error(w, "Token ausente ou inválido", http.StatusUnauthorized)
			return
		}

		// Continua para o próximo handler
		next.ServeHTTP(w, r)
	})
}
