package schemas

type CryptoInfo struct {
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"price,string"`
}
