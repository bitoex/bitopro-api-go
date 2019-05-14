package bitopro

import (
	"encoding/json"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OrderHistory struct
type OrderHistory struct {
	Data []OrderInfo `json:"data,omitempty"`
	StatusCode
}

// GetOrderHistory Ref. https://developer.bitopro.com/docs#operation/getOrderHistory
func (api *AuthAPI) GetOrderHistory() *OrderHistory {
	var data OrderHistory

	code, res := internal.ReqWithoutBody(api.identity, api.key, api.secret, "GET", "v2/orders/history")

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
