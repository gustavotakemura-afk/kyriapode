package services

import (
	"11/database"
	"11/models"
	"fmt"
)

func ListarFighters() {
	db, err := database.Conectar()
	if err != nil {
		fmt.Println("Erro conexão:", err)
		return
	}
	defer db.Close()

	query := `
	SELECT id, name, nickname, age, weight_class, wins, losses, draws, style, organization, image_url, active
	FROM fighters
	where deleted = false
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro ao buscar:", err)
		return
	}
	defer rows.Close()

	var fighters []models.Fighters

	for rows.Next() {
		var f models.Fighters

		err := rows.Scan(
			&f.ID,
			&f.Name,
			&f.Nickname,
			&f.Age,
			&f.WeightClass,
			&f.Wins,
			&f.Losses,
			&f.Draws,
			&f.Style,
			&f.Organization,
			&f.ImageURL,
			&f.Active,
		)

		if err != nil {
			fmt.Println("Erro ao ler dados:", err)
			continue
		}

		fighters = append(fighters, f)
	}

	fmt.Println("\n=== LISTA DE LUTADORES ===")

	if len(fighters) == 0 {
		fmt.Println("Nenhum lutador encontrado.")
		return
	}

	for _, f := range fighters {
		status := "Inativo"
		if f.Active {
			status = "Ativo"
		}

		fmt.Println("----------------------------")
		fmt.Printf("ID: %d\n", f.ID)
		fmt.Printf("Nome: %s (%s)\n", f.Name, f.Nickname)
		fmt.Printf("Idade: %d\n", f.Age)
		fmt.Printf("Categoria: %s\n", f.WeightClass)
		fmt.Printf("Cartel: %d-%d-%d\n", f.Wins, f.Losses, f.Draws)
		fmt.Printf("Estilo: %s\n", f.Style)
		fmt.Printf("Organização: %s\n", f.Organization)
		fmt.Printf("Status: %s\n", status)
	}
	fmt.Println("----------------------------")
}
