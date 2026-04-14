package services

import (
	"11/database"
	"fmt"
)

func ListarFightersResumo() {
	db, err := database.Conectar()
	if err != nil {
		fmt.Println("Erro conexão:", err)
		return
	}
	defer db.Close()

	query := `
	SELECT id, name, nickname
	FROM fighters
	WHERE deleted = false
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro ao buscar:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== LUTADORES DISPONÍVEIS ===")

	for rows.Next() {
		var id int
		var name, nickname string

		err := rows.Scan(&id, &name, &nickname)
		if err != nil {
			fmt.Println("Erro ao ler:", err)
			continue
		}

		fmt.Printf("ID: %d | Nome: %s (%s)\n", id, name, nickname)
	}
}
