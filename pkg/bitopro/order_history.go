package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OrderHistory struct
type OrderHistory struct {
	Data []OrderInfo `json:"data,omitempty"`
	StatusCode
}

// GetOrderHistory Ref. https://developer.bitopro.com/docs#operation/getOrderHistory
func (api *AuthAPI) GetOrderHistory(pair string) *OrderHistory {
	var data OrderHistory

	code, res := internal.ReqWithoutBody(api.identity, api.Key, api.secret, "GET", fmt.Sprintf("%s%s", "v3/orders/all/", pair))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
