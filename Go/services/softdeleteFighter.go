package services

import (
	"11/database"
	"fmt"
)

func InativarFighters(id int) {
	db, err := database.Conectar()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}

	defer db.Close()

	query := "UPDATE fighters SET deleted = true WHERE id = ?"
	_, err = db.Exec(query, id)

	if err != nil {
		fmt.Println("Erro ao inativar o lutador:", err)
		return
	}
	fmt.Println("Lutador inativado com sucesso!")

}
