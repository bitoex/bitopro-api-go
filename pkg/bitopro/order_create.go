package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// CreateOrder struct
type CreateOrder struct {
	OrderID   string `json:"orderId,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Action    string `json:"action,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Price     string `json:"price,omitempty"`
	StatusCode
}

func createOrder(api *AuthAPI, pair string, body map[string]interface{}) *CreateOrder {
	var data CreateOrder

	code, res := internal.ReqWithBody(api.identity, api.key, api.secret, fmt.Sprintf("%s/%s", "v2/orders", pair), body)

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// CreateOrderLimitBuy Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderLimitBuy(pair, price, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "limit",
		"action":    "buy",
		"price":     price,
		"amount":    amount,
		"timestamp": internal.GetTimestamp(),
	})
}

// CreateOrderLimitSell Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderLimitSell(pair, price, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "limit",
		"action":    "sell",
		"price":     price,
		"amount":    amount,
		"timestamp": internal.GetTimestamp(),
	})
}

// CreateOrderMarketBuy Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderMarketBuy(pair, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "market",
		"action":    "buy",
		"amount":    amount,
		"timestamp": internal.GetTimestamp(),
	})
}

// CreateOrderMarketSell Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderMarketSell(pair, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "market",
		"action":    "sell",
		"amount":    amount,
		"timestamp": internal.GetTimestamp(),
	})
}
