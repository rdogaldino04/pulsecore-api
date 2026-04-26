package main

import (
	"pulsecore-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler.RegisterPessoaRoutes(router)

	router.Run(":8080")
}
