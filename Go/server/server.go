package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"11/services"
)

func main() {
	// Rota que retorna todos os lutadores
	http.HandleFunc("/fighters", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		fighters, err := services.ListarFightersAPI()
		if err != nil {
			http.Error(w, "Erro ao listar lutadores", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fighters)
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
