package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruxeiro/go_api_gin/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz:":"E ai eu, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	fmt.Println(string(respostaBody))
	fmt.Println(mockDaResposta)

}
