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
type GoogleSearchResult struct {
	Items []struct {
		Title   string `json:"title"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"items"`
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configura os cabeçalhos de CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Em produção, substitua '*' pelo domínio específico
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Verifica se é uma requisição preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK) // Responde positivamente às requisições preflight
			return
		}

		// Procede para o próximo handler
		next.ServeHTTP(w, r)
	}
}

/*func buscarNoGoogleESalvar(query string) ([]Site, error) {
	//var googleAPIKey = "bruh"
	var googleCSEID = "bruh"
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s",
		googleAPIKey, googleCSEID, query)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao fazer a requisição para a Google Custom Search: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler a resposta: %v", err)
		return nil, err
	}

	var searchResults GoogleSearchResult

	err = json.Unmarshal(body, &searchResults)
	if err != nil {
		log.Printf("Erro ao fazer unmarshal dos resultados da busca: %v", err)
		return nil, err
	}

	var sites []Site
	for _, item := range searchResults.Items {
		site := Site{
			Title:       item.Title,
			Description: item.Snippet,
			URL:         item.Link,
		}
		sites = append(sites, site)

		// Inserção no banco de dados
		_, err := db.Exec("INSERT INTO sites (titulo, descricao, link) VALUES ($1, $2, $3)",
			item.Title, item.Snippet, item.Link)
		if err != nil {
			log.Printf("Erro ao inserir resultado no banco de dados: %v", err)
			// Decida como lidar com erros aqui - pode optar por continuar ou retornar um erro
		}
	}

	return sites, nil
}*/ /*
func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "A consulta de pesquisa é necessária", http.StatusBadRequest)
		return
	}

	sites, err := buscarNoGoogleESalvar(query)
	if err != nil {
		http.Error(w, "Erro ao buscar no Google", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sites)
}
*/
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Permitir apenas requisições de localhost:3000
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
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

// Handler para atualizar a chave API de um usuário
func updateApiKey(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Login  string `json:"login"`  // Supondo que o login do usuário seja enviado junto
		ApiKey string `json:"apiKey"` // A chave API a ser atualizada
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	log.Printf("Recebido update da chave API para o usuário: %s", requestBody.Login)

	var userExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM usuarios WHERE login = $1)", requestBody.Login).Scan(&userExists)
	if err != nil || !userExists {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	if userExists {
		// Executa a atualização da chave API para o usuário especificado
		_, err = db.Exec("UPDATE usuarios SET api_key = $2 WHERE login = $1", requestBody.Login, requestBody.ApiKey)
		if err != nil {
			// Se ocorrer um erro durante a atualização, retorna um erro 500
			log.Printf("Erro ao atualizar a chave API: %v", err)
			http.Error(w, "Erro interno ao atualizar a chave API", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Chave API atualizada com sucesso"})
}

func main() {
	initDB()
	log.Println("Inicializando servidor...")

	http.HandleFunc("/login", corsMiddleware(login))
	http.HandleFunc("/search", corsMiddleware(search))
	http.HandleFunc("/update-api-key", corsMiddleware(updateApiKey))

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
