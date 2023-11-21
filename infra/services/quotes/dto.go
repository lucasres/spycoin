package quotes

type quoteResponse struct {
	SymbolID       string    `json:"symbol_id"`
	BidPrice       float64   `json:"bid_price"`
	LastTradePrice lastTrade `json:"last_trade"`
}

type lastTrade struct {
	Price   float64 `json:"price"`
	TradeAt string  `json:"time_exchange"`
}
