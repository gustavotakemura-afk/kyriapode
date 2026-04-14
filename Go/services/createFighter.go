package services

import (
	"11/database"
	"11/models"
	"fmt"
)

func CriarFighter(f models.Fighters) {
	db, err := database.Conectar()
	if err != nil {
		fmt.Println("Erro de conexão", err)
		return
	}

	defer db.Close()

	query := `
	INSERT INTO fighters 
	(name, nickname, age, weight_class, wins, losses, draws, style, organization, image_url, active)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(query,
		f.Name,
		f.Nickname,
		f.Age,
		f.WeightClass,
		f.Wins,
		f.Losses,
		f.Draws,
		f.Style,
		f.Organization,
		f.ImageURL,
		f.Active,
	)
	if err != nil {
		fmt.Println("Erro ao inserir lutador", err)
		return
	}

	fmt.Println("Lutador criado com sucesso!")
}
