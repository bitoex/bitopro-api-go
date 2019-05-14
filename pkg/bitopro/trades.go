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

func getTrades(pair string) *Trade {
	var data Trade

	code, res := internal.ReqPublic(fmt.Sprintf("%s/%s", "v2/trades", pair))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetTrades Ref. https://developer.bitopro.com/docs#operation/getPairTrades
func (*PubAPI) GetTrades(pair string) *Trade {
	return getTrades(pair)
}

// GetTrades Ref. https://developer.bitopro.com/docs#operation/getPairTrades
func (*AuthAPI) GetTrades(pair string) *Trade {
	return getTrades(pair)
}
