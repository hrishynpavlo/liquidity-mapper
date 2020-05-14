package liquidity

// Liquidance model
type Liquidance struct {
	Usd float32 `json:"usd" binding:"required"`
	Eur float32 `json:"eur" binding:"required"`
	Btc float32 `json:"btc" binding:"required"`
}
