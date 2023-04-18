package models

type Aluno struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"RG"`
}

var Alunos []Aluno
