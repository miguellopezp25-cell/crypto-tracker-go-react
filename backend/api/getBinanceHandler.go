package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miguellopezp25-cell/crypto-tracker-go-react/backend/service"
)

var coinNames = map[string]string{
	"BTCUSDT":  "Bitcoin",
	"ETHUSDT":  "Ethereum",
	"BNBUSDT":  "Binance Coin",
	"XRPUSDT":  "Ripple",
	"ADAUSDT":  "Cardano",
	"SOLUSDT":  "Solana",
	"DOTUSDT":  "Polkadot",
	"DOGEUSDT": "Dogecoin",
}

func getBinanceHandler(c *gin.Context) {
	symbol := c.Param("symbol")
	symbol = strings.ToUpper(symbol)

	priceTracking, err := service.GetCryptoPriceBTC(symbol)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	name, exists := coinNames[priceTracking.Symbol]
	if !exists {
		name = priceTracking.Symbol
	}

	c.JSON(http.StatusOK, gin.H{
		"symbol": priceTracking.Symbol,
		"name":   name,
		"price":  priceTracking.Price,
	})
}
