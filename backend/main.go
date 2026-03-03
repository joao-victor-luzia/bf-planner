package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Task define nossa estrutura
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// --- BANCO DE DADOS EM MEMÓRIA ---
var (
	tasks      = []Task{}      // Nossa fatia (slice) que guardará as tarefas
	tasksMutex sync.Mutex      // O "cadeado" para evitar conflitos de escrita
	nextID     = 1             // Contador para gerar IDs automáticos
)

func main() {
	// Rota simples de teste
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Lida com o Preflight do navegador
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPost {
			var newTask Task

			err := json.NewDecoder(r.Body).Decode(&newTask)
			if (err != nil){
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}

			tasksMutex.Lock()
				newTask.ID = nextID
				nextID++
				tasks = append(tasks, newTask)
			tasksMutex.Unlock()
			
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTask)
			return;
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	fmt.Println("Backend rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}