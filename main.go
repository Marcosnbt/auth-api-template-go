// Início da estrutura do projeto - abril de 2025
// Entry point of the application
package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-api-template-go/config"
	"auth-api-template-go/handlers"
)

// Handler principal implementado em maio de 2025
func main() {
	config.LoadEnv() // Carrega variáveis de ambiente

	http.HandleFunc("/", handlers.HomeHandler)

	fmt.Println("Servidor iniciado na porta 8080")
	// Últimos testes antes da entrega - junho de 2025
	log.Fatal(http.ListenAndServe(":8080", nil))
}
