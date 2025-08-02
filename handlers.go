// Handles login and registration logic
package main
[02:52, 02/08/2025] Dani Motta: }
[02:53, 02/08/2025] Dani Motta: package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Bem-vindo à API"}
	json.NewEncoder(w).Encode(resp)
