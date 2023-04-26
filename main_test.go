package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruxeiro/go_api_gin/controller"
	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCodeSaudacaoParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/eu", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	}
}
