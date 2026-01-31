package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguellopezp25-cell/crypto-tracker-go-react/backend/config"
)

func StartServer() {
	err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	port := config.GetEnv("PORT")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":" + port) // escucha en 0.0.0.0:8081 por defecto
	fmt.Println("Server running on port " + port)
}
