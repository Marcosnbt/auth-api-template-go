// Arquivo criado em maio de 2025
// Definição das rotas da aplicação
package routes

import (
	"fmt"
	"net/http"
)

// InitRoutes registra as rotas da API
func InitRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API funcionando corretamente")
	})

	// Outras rotas podem ser adicionadas aqui
}
