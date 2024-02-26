package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	_ "github.com/lib/pq"
)

type Topic struct {
	Index       int    `json:"index"`
	Description string `json:"description"`
}

/*
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
*/

func loadAllSites() ([]string, error) {
	var documents []string
	rows, err := db.Query("SELECT titulo, descricao FROM sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var title, description string
		if err := rows.Scan(&title, &description); err != nil {
			return nil, err
		}
		// Concatena título e descrição com um espaço entre eles
		document := title + " " + description
		documents = append(documents, document)
	}
	return documents, nil
}
func performLDAPython(documents []string) ([]Topic, error) {
	documentsJSON, _ := json.Marshal(documents)

	cmd := exec.Command("python", "lda.py")
	cmd.Stdin = strings.NewReader(string(documentsJSON))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// A saída é esperada como uma matriz de arrays contendo um número e uma string.
	var rawTopics [][]interface{}
	err = json.Unmarshal(output, &rawTopics)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do JSON: %v, saída: %s", err, output)
	}

	// Converter rawTopics para []Topic
	var topics []Topic
	for _, t := range rawTopics {
		index, ok := t[0].(float64) // JSON numbers are always floats
		if !ok {
			return nil, fmt.Errorf("erro ao converter índice do tópico para int")
		}
		description, ok := t[1].(string)
		if !ok {
			return nil, fmt.Errorf("erro ao converter descrição do tópico para string")
		}
		topics = append(topics, Topic{Index: int(index), Description: description})
	}

	return topics, nil
}
