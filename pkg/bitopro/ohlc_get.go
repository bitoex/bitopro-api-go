package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OhlcInfo struct
type OhlcInfo struct {
	Timestamp int64  `json:"timestamp"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
}

// Ohlc struct
type Ohlc struct {
	Data []OhlcInfo `json:"data,omitempty"`
	StatusCode
}

func getOhlc(pair string, resolution string, from int64, to int64, proxy string) *Ohlc {
	var data Ohlc

	code, res, err := internal.ReqPublic(fmt.Sprintf("%s/%s?resolution=%s&from=%d&to=%d", "v3/trading-history", pair, resolution, from, to), proxy)
	if err != nil {
		data.Error = err.Error()
	}
	// fmt.Println(res)
	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

func (p *PubAPI) GetOhlc(pair string, resolution string, from int64, to int64) *Ohlc {
	return getOhlc(pair, resolution, from, to, p.proxy)
}

func (a *AuthAPI) GetOhlc(pair string, resolution string, from int64, to int64) *Ohlc {
	return getOhlc(pair, resolution, from, to, a.proxy)
}
