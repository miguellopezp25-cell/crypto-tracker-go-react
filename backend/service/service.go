package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CryptoInfo struct {
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"price,string"`
}

func GetCryptoPriceBTC(symbol string) (CryptoInfo, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	response, err := http.Get(url)
	if err != nil {
		return CryptoInfo{}, err
	}
	defer response.Body.Close()

	// Validar si Binance respondió con error (ej: símbolo no encontrado)
	if response.StatusCode != http.StatusOK {
		return CryptoInfo{}, errors.New("símbolo no encontrado en Binance")
	}

	responseInfo, err := io.ReadAll(response.Body)
	if err != nil {
		return CryptoInfo{}, err
	}

	var cryptoInfo CryptoInfo
	err = json.Unmarshal(responseInfo, &cryptoInfo)
	if err != nil {
		return CryptoInfo{}, err
	}
	return cryptoInfo, nil
}
