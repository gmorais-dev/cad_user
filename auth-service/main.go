package main

import (
	"auth-service/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API rodando com sucesso!"})
	})

	r.Run() // por padrão, roda na porta :8080
}
