// Handles login and registration logic
package handlers

import (
	"encoding/json"
	"net/http"
)

// HomeHandler responde com uma mensagem de boas-vindas
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"message": "Bem-vindo à API de autenticação",
	}
	json.NewEncoder(w).Encode(resp)
}

