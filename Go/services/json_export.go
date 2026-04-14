package services

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExportarFightersJSON() {
	fighters, err := ListarFightersAPI()
	if err != nil {
		fmt.Println("Erro ao listar lutadores:", err)
		return
	}

	data, err := json.MarshalIndent(fighters, "", "  ")
	if err != nil {
		fmt.Println("Erro ao gerar JSON:", err)
		return
	}

	err = os.WriteFile("fighters.json", data, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar JSON:", err)
		return
	}

	fmt.Println("Arquivo fighters.json gerado com sucesso!")
}
