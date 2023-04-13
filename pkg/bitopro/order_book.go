package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OrderBookInfo struct
type OrderBookInfo struct {
	Amount string `json:"amount"`
	Price  string `json:"price"`
	Count  int    `json:"count"`
	Total  string `json:"total"`
}

// OrderBook struct
type OrderBook struct {
	Bids []OrderBookInfo `json:"bids"`
	Asks []OrderBookInfo `json:"asks"`
	StatusCode
}

func getOrderBook(pair string, limit int, proxy string) *OrderBook {
	var data OrderBook

	code, res, err := internal.ReqPublic(fmt.Sprintf("%s/%s?limit=%d", "v3/order-book", pair, limit), proxy)
	if err != nil {
		data.Error = err.Error()
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetOrderBook Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (p *PubAPI) GetOrderBook(pair string) *OrderBook {
	return getOrderBook(pair, 5, p.proxy)
}

// GetOrderBook Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (a *AuthAPI) GetOrderBook(pair string) *OrderBook {
	return getOrderBook(pair, 5, a.proxy)
}

// GetOrderBookWithLimit Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (p *PubAPI) GetOrderBookWithLimit(pair string, limit int) *OrderBook {
	return getOrderBook(pair, limit, p.proxy)
}

// GetOrderBookWithLimit Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (a *AuthAPI) GetOrderBookWithLimit(pair string, limit int) *OrderBook {
	return getOrderBook(pair, limit, a.proxy)
}
