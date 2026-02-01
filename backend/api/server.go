package api

import (
	"fmt"

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

	router.GET("/binance/tracking/:symbol", getBinanceHandler)

	router.Run(":" + port) // escucha en 0.0.0.0:8081 por defecto
	fmt.Println("Server running on port " + port)
}
