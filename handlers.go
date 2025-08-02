// Handles login and registration logic
package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Bem-vindo à API"}
	json.NewEncoder(w).Encode(resp)
}

