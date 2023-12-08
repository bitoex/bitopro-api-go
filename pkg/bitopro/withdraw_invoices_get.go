package bitopro

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/bitoex/bitopro-api-go/internal"
)

// WithdrawHistory struct
type WithdrawHistory struct {
	Data []TransactionInfo `json:"data,omitempty"`
	StatusCode
}

// WithdrawOrder struct
type WithdrawOrder struct {
	Data TransactionInfo `json:"data,omitempty"`
	StatusCode
}

// GetWithdrawHistory
func (api *AuthAPI) GetWithdrawHistory(currency string, queryParams url.Values) *WithdrawHistory {
	var data WithdrawHistory

	queryStr := queryParams.Encode()
	endpoint := fmt.Sprintf("%s/%s", "v3/wallet/withdrawHistory", currency)
	if queryStr != "" {
		endpoint = fmt.Sprintf("%s/%s?%s", "v3/wallet/withdrawHistory", currency, queryStr)
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

func (api *AuthAPI) GetWithdrawBySerial(currency string, serial string) *WithdrawOrder {
	var data WithdrawOrder

	endpoint := fmt.Sprintf("%s/%s/%s", "v3/wallet/withdraw", currency, serial)
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

func (api *AuthAPI) GetWithdrawById(currency string, id string) *WithdrawOrder {
	var data WithdrawOrder

	endpoint := fmt.Sprintf("%s/%s/id/%s", "v3/wallet/withdraw", currency, id)
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
