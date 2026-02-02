package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miguellopezp25-cell/crypto-tracker-go-react/backend/config"
)

func StartServer() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	gin.SetMode(cfg.Server.Mode)

	port := cfg.Server.Port
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	router.GET("/binance/tracking/:symbol", getBinanceHandler)

	err = router.Run(":" + port) // escucha en 0.0.0.0:8081 por defecto
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server running on port " + port)

}
