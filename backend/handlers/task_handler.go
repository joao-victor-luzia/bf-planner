package handlers

import (
    "encoding/json"
    "net/http"
    "sync"
    "time"
    "bf-planner/server/models" // Importando do seu próprio módulo
)

var (
	tasks      = []models.Task{}      
	tasksMutex sync.Mutex      
	nextID     = 1             
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
    // Configuração de CORS (Global para este handler)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method == http.MethodPost {
        var newTask models.Task
        if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
            http.Error(w, "JSON inválido", http.StatusBadRequest)
            return
        }

        tasksMutex.Lock()
        newTask.ID = nextID
        newTask.TimCreated = time.Now().UTC()
        nextID++
        tasks = append(tasks, newTask)
        tasksMutex.Unlock()

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newTask)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}