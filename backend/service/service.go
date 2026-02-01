package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	schema "github.com/miguellopezp25-cell/crypto-tracker-go-react/backend/schemas"
)

func GetCryptoPriceBTC(symbol string) (schema.CryptoInfo, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	response, err := http.Get(url)
	if err != nil {
		return schema.CryptoInfo{}, err
	}
	defer response.Body.Close()

	// Validar si Binance respondió con error (ej: símbolo no encontrado)
	if response.StatusCode != http.StatusOK {
		return schema.CryptoInfo{}, errors.New("símbolo no encontrado en Binance")
	}

	responseInfo, err := io.ReadAll(response.Body)
	if err != nil {
		return schema.CryptoInfo{}, err
	}

	var cryptoInfo schema.CryptoInfo
	err = json.Unmarshal(responseInfo, &cryptoInfo)
	if err != nil {
		return schema.CryptoInfo{}, err
	}
	return cryptoInfo, nil
}
