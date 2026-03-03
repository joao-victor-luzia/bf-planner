package main

import (
    "fmt"
    "net/http"
    "bf-planner/server/handlers" 
)

func main() {
    http.HandleFunc("/api/tasks", handlers.TaskHandler)

    port := ":8080"
    fmt.Printf("Backend rodando em http://localhost%s\n", port)
    
    err := http.ListenAndServe(port, nil)
    if err != nil {
        fmt.Printf("Erro ao iniciar servidor: %v\n", err)
    }
}