// Arquivo criado em maio
package routes

import (
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API OK"))
	})
}
