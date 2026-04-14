package main

import (
	"fmt"
	"strconv"

	"11/models"
	"11/services"
	"11/utils"
)

func createFighter() {
	var f models.Fighters
	fmt.Println("\n=== Cadastro de Lutador ===")

	f.Name = utils.LerString("Nome: ")

	f.Nickname = utils.LerString("Apelido: ")

	idadeStr := utils.LerString("Idade: ")
	f.Age, _ = strconv.Atoi(idadeStr)

	f.WeightClass = utils.LerString("Categoria de Peso: ")

	vitoriasStr := utils.LerString("Vitorias: ")
	f.Wins, _ = strconv.Atoi(vitoriasStr)

	derrotasStr := utils.LerString("Derrotas: ")
	f.Losses, _ = strconv.Atoi(derrotasStr)

	empatesStr := utils.LerString("Empates: ")
	f.Draws, _ = strconv.Atoi(empatesStr)

	f.Style = utils.LerString("Estilo de Luta: ")

	f.Organization = utils.LerString("Organização de Luta: ")

	f.ImageURL = utils.LerString("URL da Imagem do Lutador: ")

	ativoStr := utils.LerString("O lutador está ativo? (s/n): ")
	f.Active = ativoStr == "s" || ativoStr == "S"

	services.CriarFighter(f)
}

func atualizarFighter() {
	var f models.Fighters
	services.ListarFightersResumo()

	fmt.Println("\n=== Atualizar Lutador ===")

	idStr := utils.LerString("ID do lutador: ")
	f.ID, _ = strconv.Atoi(idStr)

	f.Name = utils.LerString("Novo nome: ")
	f.Nickname = utils.LerString("Novo apelido: ")

	idadeStr := utils.LerString("Nova idade: ")
	f.Age, _ = strconv.Atoi(idadeStr)

	f.WeightClass = utils.LerString("Nova categoria: ")

	winsStr := utils.LerString("Vitórias: ")
	f.Wins, _ = strconv.Atoi(winsStr)

	lossStr := utils.LerString("Derrotas: ")
	f.Losses, _ = strconv.Atoi(lossStr)

	drawStr := utils.LerString("Empates: ")
	f.Draws, _ = strconv.Atoi(drawStr)

	f.Style = utils.LerString("Estilo: ")
	f.Organization = utils.LerString("Organização: ")
	f.ImageURL = utils.LerString("URL da imagem: ")

	ativoStr := utils.LerString("Ativo (s/n): ")
	f.Active = ativoStr == "s" || ativoStr == "S"

	services.AtualizarFighter(f)
}

func InativarFighter() {
	services.ListarFightersResumo()

	fmt.Println("\n=== Inativar Lutador ===")
	idStr := utils.LerString("ID do lutador a inativar: ")
	id, _ := strconv.Atoi(idStr)
	services.InativarFighters(id)
}
func main() {

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Cadastrar Lutador")
		fmt.Println("2. Listar Lutadores")
		fmt.Println("3. Atualizar Lutador")
		fmt.Println("4. Inativar Lutador (soft delete)")
		fmt.Println("5. Exportar Lutadores para JSON")
		fmt.Println("6. Sair")

		opcaoStr := utils.LerString("Escolha uma opção: ")
		opcao, _ := strconv.Atoi(opcaoStr)

		switch opcao {
		case 1:
			createFighter()
		case 2:
			services.ListarFighters()
		case 3:
			atualizarFighter()
		case 4:
			InativarFighter()
		case 5:
			services.ExportarFightersJSON()

		case 6:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}
