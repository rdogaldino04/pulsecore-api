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
	result, err := service.ListarPessoas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Success: true,
		Data:    result,
	})
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

	pessoa, err := service.CriarPessoa(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ApiResponse{
		Success: true,
		Data:    pessoa,
	})
}
