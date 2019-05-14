package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// CancelOrder struct
type CancelOrder struct {
	OrderID   string `json:"orderId,omitempty"`
	Action    string `json:"action,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Price     string `json:"price,omitempty"`
	Amount    string `json:"amount,omitempty"`
	StatusCode
}

// CancelOrder Ref. https://developer.bitopro.com/docs#operation/cancelOrder
func (api *AuthAPI) CancelOrder(pair string, orderID int) *CancelOrder {
	var data CancelOrder

	code, res := internal.ReqWithoutBody(api.identity, api.key, api.secret, "DELETE", fmt.Sprintf("%s/%s/%d", "v2/orders", pair, orderID))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
