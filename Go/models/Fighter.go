package models

type Fighters struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`     // apelido
	Age          int    `json:"age"`          // idade
	WeightClass  string `json:"weight_class"` // categoria de peso
	Wins         int    `json:"wins"`         // vitorias
	Losses       int    `json:"losses"`       // derrotas
	Draws        int    `json:"draws"`        // empates
	Style        string `json:"style"`        // estilo de luta
	Organization string `json:"organization"` // organização de luta
	ImageURL     string `json:"image_url"`    // url da imagem do lutador
	Active       bool   `json:"active"`       // se o lutador está ativo ou aposentado
	deleted      bool   `json:"deleted"`      // para marcar o lutador como deletado (soft delete)
}
