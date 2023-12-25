package bitopro

import (
	"encoding/json"

	"github.com/bitoex/bitopro-api-go/internal"
)

type TradingPairInfo struct {
    Pair                        string  `json:"pair"`
    Base                        string  `json:"base"`
    Quote                       string  `json:"quote"`
    BasePrecision               string  `json:"basePrecision"`
    QuotePrecision              string  `json:"quotePrecision"`
    MinLimitBaseAmount          string  `json:"minLimitBaseAmount"`
    MaxLimitBaseAmount          string  `json:"maxLimitBaseAmount"`
    MinMarketBuyQuoteAmount     string  `json:"minMarketBuyQuoteAmount"`
    OrderOpenLimit              string  `json:"orderOpenLimit"`
    Maintain                    bool    `json:"maintain"`
    OrderBookQuotePrecision     string  `json:"orderBookQuotePrecision"`
    OrderBookQuoteScaleLevel    string  `json:"orderBookQuoteScaleLevel"`
    AmountPrecision             string  `json:"amountPrecision"`
}

type TradingPairInfos struct {
    Data []TradingPairInfo `json:"data,omitempty"`
    StatusCode
}

func getTradingPairInfos(proxy string) *TradingPairInfos {
	var data TradingPairInfos

	code, res, err := internal.ReqPublic("v3/provisioning/trading-pairs", proxy)
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


// GetTradingPairInfos Ref. https://github.com/bitoex/bitopro-offical-api-docs/blob/master/api/v3/public/get_trading_pair_info.md
func (p *PubAPI) GetTradingPairInfos() *TradingPairInfos {
	return getTradingPairInfos(p.proxy)
}

// GetTradingPairInfos Ref. https://github.com/bitoex/bitopro-offical-api-docs/blob/master/api/v3/public/get_trading_pair_info.md
func (a *AuthAPI) GetTradingPairInfos() *TradingPairInfos {
	return getTradingPairInfos(a.proxy)
}
