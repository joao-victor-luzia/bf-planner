package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Task define como é uma tarefa no nosso sistema
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	// Rota simples de teste
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		// Permitir que o Svelte acesse o Go (CORS)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		tasks := []Task{
			{ID: 1, Title: "Aprender Go", Done: false},
			{ID: 2, Title: "Configurar Svelte", Done: true},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	fmt.Println("Backend rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}