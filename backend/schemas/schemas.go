package schemas

//struc to hold crypto information
type CryptoInfo struct {
	ID     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Change float64 `json:"change_24h"`
}
