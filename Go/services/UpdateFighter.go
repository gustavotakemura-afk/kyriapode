package services

import (
	"11/database"
	"11/models"
	"fmt"
)

func AtualizarFighter(f models.Fighters) {
	db, err := database.Conectar()
	if err != nil {
		fmt.Println("Erro conexão:", err)
		return
	}
	defer db.Close()

	query := `
	UPDATE fighters SET 
	name=?, nickname=?, age=?, weight_class=?, 
	wins=?, losses=?, draws=?, style=?, organization=?, image_url=?, active=?
	WHERE id=?`

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
		f.ID,
	)

	if err != nil {
		fmt.Println("Erro ao atualizar:", err)
		return
	}

	fmt.Println("Lutador atualizado com sucesso!")
}
