package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// TickerInfo struct
type TickerInfo struct {
	Pair            string `json:"pair"`
	IsBuyer         bool   `json:"isBuyer"`
	High24hr        string `json:"high24hr"`
	Low24hr         string `json:"low24hr"`
	PriceChange24hr string `json:"priceChange24hr"`
	Volume24hr      string `json:"volume24hr"`
	LastPrice       string `json:"lastPrice"`
}

// Ticker struct
type Ticker struct {
	Data TickerInfo `json:"data,omitempty"`
	StatusCode
}

func getTicker(pair, proxy string) *Ticker {
	var data Ticker

	code, res := internal.ReqPublic(fmt.Sprintf("%s/%s", "v3/tickers", pair), proxy)

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetTicker Ref. https://developer.bitopro.com/docs#operation/getTickers
func (p *PubAPI) GetTicker(pair string) *Ticker {
	return getTicker(pair, p.proxy)
}

// GetTicker Ref. https://developer.bitopro.com/docs#operation/getTickers
func (a *AuthAPI) GetTicker(pair string) *Ticker {
	return getTicker(pair, a.proxy)
}
