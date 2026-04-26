package handler

import (
	"net/http"
	"pulsecore-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterPessoaRoutes(r *gin.Engine) {
	r.GET("/pessoas", getPessoas)
	r.POST("/pessoas", createPessoa) // Implementação futura
}

func getPessoas(c *gin.Context) {
	result := service.ListarPessoas()
	c.JSON(http.StatusOK, result)
}

func createPessoa(c *gin.Context) {
	var req service.CreatePessoaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	pessoa := service.CriarPessoa(req)

	c.JSON(http.StatusCreated, ApiResponse{
		Success: true,
		Data:    pessoa,
	})
}
