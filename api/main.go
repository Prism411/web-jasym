package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estrutura para os dados de login
type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Permitir apenas requisições de localhost:3000
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")   // Métodos permitidos
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("Requisição de login recebida")
	enableCors(&w) // Habilita o CORS

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK) // Responde positivamente às requisições OPTIONS
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var data LoginData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Apenas loga o login e senha se a decodificação foi bem-sucedida
	log.Printf("Tentativa de login: Login: %s, Senha: %s", data.Login, data.Password)

	// Envia uma resposta de sucesso
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"message\": \"Login bem-sucedido para: %s\"}", data.Login)
}

func search(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) // Habilita o CORS

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK) // Responde positivamente às requisições OPTIONS
		return
	}

	query := r.URL.Query().Get("query")
	log.Printf("Search received: %s\n", query)

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(map[string]string{"message": fmt.Sprintf("Search results for: %s", query)})
	if err != nil {
		http.Error(w, "Erro ao criar resposta JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func main() {
	log.Println("Inicializando servidor...")

	http.HandleFunc("/login", login)
	http.HandleFunc("/search", search)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
