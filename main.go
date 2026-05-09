package main

import (
	"pulsecore-api/internal/config"
	"pulsecore-api/internal/handler"
	"pulsecore-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	repository.InitDB()

	router := gin.Default()

	handler.RegisterPessoaRoutes(router)

	router.Run(":8080")
}
