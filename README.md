# Auth API Template (Go)

Simple JWT authentication template in Go.

## Histórico de Desenvolvimento

- *Abril 2025*: Estrutura inicial do projeto (handlers, middlewares, conexão com banco de dados).
- *Maio 2025*: Implementação de autenticação via token e rotas protegidas.
- *Junho 2025*: Refatoração modular e testes manuais nos principais endpoints.
- *Julho 2025*: Finalização e preparação para entrega ao cliente.

---

## 🔧 Tecnologias Utilizadas

- Go 1.18+
- JWT para autenticação
- Módulo net/http para rotas
- Estrutura modular por pacotes
- Carregamento de variáveis com godotenv

---

## 📁 Estrutura do Projeto

auth-api-template-go/
├── config/             # Carregamento de variáveis de ambiente (.env)
├── handlers/           # Lógica de login, registro e responses
├── middlewares/        # Validação JWT e controle de acesso
├── routes/             # Definição das rotas e health check
├── services/           # Simulação de processamento interno
├── simulacao.go        # Arquivo adicional para manter estrutura de testes
├── main.go             # Inicialização do servidor
└── go.mod              # Definições do módulo

---

## 🧪 Testes e Execução

```bash
go mod tidy
go run main.go

Acesse via: http://localhost:8080/health para verificar o status da API.

⸻

🔐 Observações

Este projeto foi desenvolvido sob demanda com entregas parciais entre abril e julho. Algumas partes foram mantidas privadas a pedido do cliente.
