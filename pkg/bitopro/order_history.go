package bitopro

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OrderHistory struct
type OrderHistory struct {
	Data []OrderInfo `json:"data,omitempty"`
	StatusCode
}

// GetOrderHistory Ref. https://developer.bitopro.com/docs#operation/getOrderHistory
func (api *AuthAPI) GetOrderHistory(pair string, queryParams url.Values) *OrderHistory {
	var data OrderHistory

	queryStr := queryParams.Encode()
	endpoint := fmt.Sprintf("%s%s", "v3/orders/all/", pair)
	if queryStr != "" {
		endpoint = fmt.Sprintf("%s%s?%s", "v3/orders/all/", pair, &queryStr)
	}
	code, res, err := internal.ReqWithoutBody(api.identity, api.Key, api.secret, "GET", endpoint, api.proxy)

	if err != nil {
		data.Error = fmt.Sprintf("req err:[%+v], res:[%+v]", err, res)
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
