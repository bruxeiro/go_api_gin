package main

import (
	"github.com/bruxeiro/go_api_gin/database"
	"github.com/bruxeiro/go_api_gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()

}
