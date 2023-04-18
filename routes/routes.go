package routes

import (
	"github.com/bruxeiro/go_api_gin/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	//Definindo uma rota, neste caso a rota Dedault para aprender outros conceitos antes
	r := gin.Default()
	//Utilizando a rota em conjunto com um modelo de requisição HTTP, nesse caso o GET
	//Definindo um Endpoint: "/alunos" e passando uma função que devolve os dados em JSON usando o Gin
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.Run()
}
