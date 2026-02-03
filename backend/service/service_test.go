package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

func TestGetCryptoPriceBTC(t *testing.T) {
	tests := []struct {
		name          string
		symbol        string
		mockResponse  string
		mockStatus    int
		expectedPrice float64
		expectError   bool
	}{
		{
			name:          "Valid symbol BTCUSDT",
			symbol:        "BTCUSDT",
			mockResponse:  `{"symbol":"BTCUSDT","price":"30000.00"}`,
			mockStatus:    http.StatusOK,
			expectedPrice: 30000.00,
			expectError:   false,
		},
		{
			name:          "Invalid symbol",
			symbol:        "INVALID",
			mockResponse:  `{"code":-1121,"msg":"Invalid symbol"}`,
			mockStatus:    http.StatusBadRequest,
			expectedPrice: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. CREAR UN SERVIDOR MOCK (Simula Binance)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatus)
				fmt.Fprintln(w, tt.mockResponse)
			}))
			defer server.Close()

			// 2. INYECTAR LA URL DEL MOCK EN VIPER
			// Esto evita que Viper busque el archivo físico y use valores en memoria
			viper.Set("api.API_URL_CRYPTO", server.URL)

			// 3. EJECUTAR LA FUNCIÓN
			priceInfo, err := GetCryptoPriceBTC(tt.symbol)

			// 4. VALIDACIONES
			if tt.expectError {
				if err == nil {
					t.Errorf("Se esperaba un error pero no se obtuvo ninguno")
				}
			} else {
				if err != nil {
					t.Errorf("Error inesperado: %v", err)
				}
				// Verificamos que el precio coincida (asumiendo que priceInfo.Price es float64)
				if priceInfo.Price != tt.expectedPrice {
					t.Errorf("Se esperaba precio %f pero se obtuvo %f", tt.expectedPrice, priceInfo.Price)
				}
			}
		})
	}
}
