package bitopro

import (
	"encoding/json"

	"github.com/bitoex/bitopro-api-go/internal"
)

// TradingPairsInfo struct
type TradingPairsInfo struct {
	Pair                     string `json:"pair"`
	Base                     string `json:"base"`
	Quote                    string `json:"quote"`
	BasePrecision            string `json:"basePrecision"`
	QuotePrecision           string `json:"quotePrecision"`
	MinLimitBaseAmount       string `json:"minLimitBaseAmount"`
	MaxLimitBaseAmount       string `json:"maxLimitBaseAmount"`
	MinMarketBuyQuoteAmount  string `json:"minMarketBuyQuoteAmount"`
	OrderOpenLimit           string `json:"orderOpenLimit"`
	Maintain                 bool   `json:"maintain"`
	OrderBookQuotePrecision  string `json:"orderBookQuotePrecision"`
	OrderBookQuoteScaleLevel string `json:"orderBookQuoteScaleLevel"`
	AmountPrecision          string `json:"amountPrecision"`
}

// TradingPairs struct
type TradingPairs struct {
	Data []TradingPairsInfo `json:"data,omitempty"`
	StatusCode
}

func getTradingpairs(proxy string) *TradingPairs {
	var data TradingPairs

	code, res, err := internal.ReqPublic("v3/provisioning/trading-pairs", proxy)
	if err != nil {
		data.Error = err.Error()
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

func (p *PubAPI) GetTradingpairs() *TradingPairs {
	return getTradingpairs(p.proxy)
}

func (a *AuthAPI) GetTradingpairs() *TradingPairs {
	return getTradingpairs(a.proxy)
}
