package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

type CreateOrder struct {
	OrderID     string `json:"orderId,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	Action      string `json:"action,omitempty"`
	Amount      string `json:"amount,omitempty"`
	IsBuy       string `json:"isBuy,omitempty"`
	Price       string `json:"price,omitempty"`
	TimeInForce string `json:"timeInForce,omitempty"`
	ClientID    int    `json:"clientId,omitempty"`
	Type        string `json:"type,omitempty"`
	StatusCode
}

func createOrder(api *AuthAPI, pair string, body map[string]interface{}) *CreateOrder {
	var data CreateOrder

	code, res, err := internal.ReqWithBody(api.identity, api.Key, api.secret, fmt.Sprintf("%s/%s", "v3/orders", pair), api.proxy, body)

	if err != nil {
		data.Error = fmt.Sprintf("req err:[%+v], res:[%+v]", err, res)
	}

	if err = json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// CreateOrderLimitBuy Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderLimitBuy(clientID int, pair, price, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "limit",
		"action":    "buy",
		"price":     price,
		"amount":    amount,
		"clientId":  clientID,
		"timestamp": internal.GetTimestamp(),
	})
}

// CreateOrderLimitSell Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderLimitSell(clientID int, pair, price, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "limit",
		"action":    "sell",
		"price":     price,
		"amount":    amount,
		"clientId":  clientID,
		"timestamp": internal.GetTimestamp(),
	})
}

// CreateOrderMarketBuy Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderMarketBuy(clientID int, pair, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "market",
		"action":    "buy",
		"amount":    amount,
		"clientId":  clientID,
		"timestamp": internal.GetTimestamp(),
	})
}

// CreateOrderMarketSell Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) CreateOrderMarketSell(clientID int, pair, amount string) *CreateOrder {
	return createOrder(api, pair, map[string]interface{}{
		"type":      "market",
		"action":    "sell",
		"amount":    amount,
		"clientId":  clientID,
		"timestamp": internal.GetTimestamp(),
	})
}

type OrderData struct {
	Pair      string `json:"pair,omitempty"`
	Action    string `json:"action,omitempty"`
	Type      string `json:"type,omitempty"`
	Price     string `json:"price,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type BatchCreateOrders struct {
	Data []struct {
		OrderID     string `json:"orderId,omitempty"`
		Timestamp   int64  `json:"timestamp,omitempty"`
		Action      string `json:"action,omitempty"`
		Amount      string `json:"amount,omitempty"`
		IsBuy       string `json:"isBuy,omitempty"`
		Price       string `json:"price,omitempty"`
		TimeInForce string `json:"timeInForce,omitempty"`
		ClientID    int    `json:"clientId,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"data"`
	StatusCode
}

// CreateOrderMarketSell Ref. https://developer.bitopro.com/docs#operation/createOrder
func (api *AuthAPI) BatchCreateOrders(orders *OrderData) *BatchCreateOrders {
	resp := BatchCreateOrders{}
	data := map[string]interface{}{
		"data": orders,
	}

	code, res, err := internal.ReqWithBody(api.identity, api.Key, api.secret, "/v3/orders/batch", api.proxy, data)
	if err != nil {
		resp.Error = fmt.Sprintf("req err:[%+v], res:[%+v]", err, res)
	}

	if code != 200 {
		resp.Error = res
	} else {
		if err = json.Unmarshal([]byte(res), &resp); err != nil {
			resp.Error = fmt.Sprintf("unmarshal failed, err:%+v", err)
		}
	}

	return &resp
}
