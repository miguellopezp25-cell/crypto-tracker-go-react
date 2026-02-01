package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miguellopezp25-cell/crypto-tracker-go-react/backend/service"
)

func getBinanceHandler(c *gin.Context) {
	symbol := c.Param("symbol")
	symbol = strings.ToUpper(symbol)

	priceTracking, err := service.GetCryptoPriceBTC(symbol)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	name := priceTracking.Name
	if name == "" {
		if priceTracking.Symbol == "BTCUSDT" {
			name = "Bitcoin"

		}
		if priceTracking.Symbol == "ETHUSDT" {
			name = "Ethereum"

		}
		if priceTracking.Symbol == "BNBUSDT" {
			name = "Binance Coin"

		}
		if priceTracking.Symbol == "XRPUSDT" {
			name = "Ripple"

		}
		if priceTracking.Symbol == "ADAUSDT" {
			name = "Cardano"

		}
		if priceTracking.Symbol == "SOLUSDT" {
			name = "Solana"

		}

	}

	c.JSON(http.StatusOK, gin.H{
		"symbol": priceTracking.Symbol,
		"name":   name,
		"price":  priceTracking.Price,
	})
}
