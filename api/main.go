package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Site struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"` // Modificado de `Link` para `URL`
}

var db *sql.DB

// Estrutura para os dados de login
type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type SearchResult struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"` // Adicione uma URL para cada resultado
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Permitir apenas requisições de localhost:3000
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")   // Métodos permitidos
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func initDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=localhost port=5433 user=postgres dbname=postgres sslmode=disable password=admin")
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("Requisição de login recebida")
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método não permitido"})
		return
	}

	var data LoginData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erro ao decodificar JSON"})
		return
	}

	// Busca a senha hash armazenada para o login fornecido
	var senhaHashArmazenada string
	err = db.QueryRow("SELECT senha FROM usuarios WHERE login = $1", data.Login).Scan(&senhaHashArmazenada)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Usuário ou senha inválidos"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Erro ao processar a solicitação"})
			log.Printf("Erro ao buscar usuário: %v", err)
		}
		return
	}

	// Compara a senha fornecida com a hash armazenada
	err = bcrypt.CompareHashAndPassword([]byte(senhaHashArmazenada), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Usuário ou senha inválidos"})
		return
	}

	// Autenticação bem-sucedida
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login bem-sucedido"})
}

// COMEÇA AQUI SEARCH
// SEARCH ABAIXO
func search(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	rows, err := db.Query("SELECT id, titulo AS title, descricao AS description, link AS url FROM sites")
	if err != nil {
		http.Error(w, "Erro ao buscar sites", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sites []Site

	for rows.Next() {
		var s Site
		if err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.URL); err != nil {
			http.Error(w, "Erro ao ler resultado da busca", http.StatusInternalServerError)
			return
		}
		sites = append(sites, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sites)
}

func main() {
	initDB()
	log.Println("Inicializando servidor...")

	http.HandleFunc("/login", login)
	http.HandleFunc("/search", search)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
