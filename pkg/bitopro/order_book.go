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

func getOrderBook(pair string) *OrderBook {
	var data OrderBook

	code, res := internal.ReqPublic(fmt.Sprintf("%s/%s", "v2/order-book", pair))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetOrderBook Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (*PubAPI) GetOrderBook(pair string) *OrderBook {
	return getOrderBook(pair)
}

// GetOrderBook Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (*AuthAPI) GetOrderBook(pair string) *OrderBook {
	return getOrderBook(pair)
}
