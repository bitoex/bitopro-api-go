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

func getTickers(proxy string) *Tickers {
	var data Tickers

	code, res, err := internal.ReqPublic("v3/tickers", proxy)
	if err != nil {
		data.Error = err.Error()
	} else {
		if err := json.Unmarshal([]byte(res), &data); err != nil {
			data.Error = res
		}
	}

	data.Code = code

	return &data
}

// GetTickers Ref. https://developer.bitopro.com/docs#operation/getTickers
func (p *PubAPI) GetTickers() *Tickers {
	return getTickers(p.proxy)
}

// GetTickers Ref. https://developer.bitopro.com/docs#operation/getTickers
func (a *AuthAPI) GetTickers() *Tickers {
	return getTickers(a.proxy)
}
