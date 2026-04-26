package handler

import (
	"net/http"
	"pulsecore-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterPessoaRoutes(r *gin.Engine) {
	r.GET("/pessoas", getPessoas)
}

func getPessoas(c *gin.Context) {
	result := service.ListarPessoas()
	c.JSON(http.StatusOK, result)
}
