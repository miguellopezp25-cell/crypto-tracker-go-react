package service

import "testing"

func TestGetCryptoPriceBTC(t *testing.T) {

	tests := []struct {
		name          string
		symbol        string
		mockResponse  string
		expectedPrice float64
		expectError   bool
	}{
		{
			name:          "Valid symbol BTCUSDT",
			symbol:        "BTCUSDT",
			mockResponse:  `{"symbol":"BTCUSDT","price":"30000.00","name":"Bitcoin"}`,
			expectedPrice: 30000.00,
			expectError:   false,
		},
		{
			name:          "Invalid symbol",
			symbol:        "INVALID",
			mockResponse:  `{"error": "Invalid symbol"}`,
			expectedPrice: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priceInfo, err := GetCryptoPriceBTC(tt.symbol)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tt.expectedPrice <= 0 {
					t.Errorf("Expected price %f but got %f", tt.expectedPrice, priceInfo.Price)
				}
			}
		})
	}
}
