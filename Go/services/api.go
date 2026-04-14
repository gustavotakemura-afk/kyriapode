package services

import (
	"11/database"
	"11/models"
)

func ListarFightersAPI() ([]models.Fighters, error) {
	db, err := database.Conectar()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT id, name, nickname, age, weight_class, wins, losses, draws, style, organization, image_url, active
	FROM fighters
	WHERE deleted = false
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fighters []models.Fighters

	for rows.Next() {
		var f models.Fighters
		err := rows.Scan(
			&f.ID, &f.Name, &f.Nickname, &f.Age, &f.WeightClass, &f.Wins,
			&f.Losses, &f.Draws, &f.Style, &f.Organization, &f.ImageURL, &f.Active,
		)
		if err != nil {
			continue
		}
		fighters = append(fighters, f)
	}

	return fighters, nil
}
