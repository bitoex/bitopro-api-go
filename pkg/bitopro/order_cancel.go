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

	code, res, err := internal.ReqWithoutBody(api.identity, api.Key, api.secret, "DELETE", fmt.Sprintf("%s/%s/%d", "v3/orders", pair, orderID), api.proxy)

	if err != nil {
		data.Error = fmt.Sprintf("req err:[%+v], res:[%+v]", err, res)
	}

	if err = json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

type CancelAllResp struct {
	Data  map[string][]string
	Error string
	Code  int
}

func (api *AuthAPI) CancelAll(pair string) *CancelAllResp {
	var data CancelAllResp

	code, res, err := internal.ReqWithoutBody(api.identity, api.Key, api.secret, "DELETE", fmt.Sprintf("%s/%s", "v3/orders", pair), api.proxy)

	if err != nil {
		data.Error = fmt.Sprintf("req err:[%+v], res:[%+v]", err, res)
	}

	if err = json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
