package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// TradeInfo struct
type TradeInfo struct {
	Amount    string `json:"amount"`
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	IsBuyer   bool   `json:"isBuyer"`
}

// Trade struct
type Trade struct {
	Data []TradeInfo `json:"data,omitempty"`
	StatusCode
}

func getTrades(pair, proxy string) *Trade {
	var data Trade

	code, res, err := internal.ReqPublic(fmt.Sprintf("%s/%s", "v3/trades", pair), proxy)
	if err != nil {
		data.Error = err.Error()
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetTrades Ref. https://developer.bitopro.com/docs#operation/getPairTrades
func (p *PubAPI) GetTrades(pair string) *Trade {
	return getTrades(pair, p.proxy)
}

// GetTrades Ref. https://developer.bitopro.com/docs#operation/getPairTrades
func (a *AuthAPI) GetTrades(pair string) *Trade {
	return getTrades(pair, a.proxy)
}
