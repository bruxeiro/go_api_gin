package main

import (
	"github.com/bruxeiro/go_api_gin/database"
	"github.com/bruxeiro/go_api_gin/models"
	"github.com/bruxeiro/go_api_gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Bruxeiro", CPF: "00000000000", RG: "99999999999"},
		{Nome: "Bruxeira", CPF: "00000000000", RG: "99999999991"},
		{Nome: "Bruxo", CPF: "00000000000", RG: "99999999992"},
	}
	routes.HandleRequests()

}
