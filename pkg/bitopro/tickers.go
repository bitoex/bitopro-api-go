package bitopro

import (
	"encoding/json"

	"github.com/bitoex/bitopro-api-go/internal"
)

// Tickers struct
type Tickers struct {
	Data []TickerInfo `json:"data,omitempty"`
	StatusCode
}

func getTickers() *Tickers {
	var data Tickers

	code, res := internal.ReqPublic("v2/tickers")

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetTickers Ref. https://developer.bitopro.com/docs#operation/getTickers
func (*PubAPI) GetTickers() *Tickers {
	return getTickers()
}

// GetTickers Ref. https://developer.bitopro.com/docs#operation/getTickers
func (*AuthAPI) GetTickers() *Tickers {
	return getTickers()
}
