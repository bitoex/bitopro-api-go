package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

type CreateWithdrawInfo struct {
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Fee      string `json:"fee"`
	Protocol string `json:"protocol"`
	Serial   string `json:"serial"`
	Total    string `json:"total"`
}

type CreateWithdraw struct {
	Data CreateWithdrawInfo `json:"data,omitempty"`
	StatusCode
}

func createwithdraw(api *AuthAPI, currency string, body map[string]interface{}) *CreateWithdraw {
	var data CreateWithdraw

	code, res, err := internal.ReqWithBody(api.identity, api.Key, api.secret, fmt.Sprintf("%s/%s", "v3/wallet/withdraw", currency), api.proxy, body)

	if err != nil {
		data.Error = fmt.Sprintf("req err:[%+v], res:[%+v]", err, res)
	}

	if err = json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

func (api *AuthAPI) CreateWithdraw(currency, protocol, address, amount string, options ...string) *CreateWithdraw {
	var message string
	if len(options) > 0 {
		message = options[0]
	} else {
		message = ""
	}
	return createwithdraw(api, currency, map[string]interface{}{
		"protocol":  protocol,
		"address":   address,
		"amount":    amount,
		"message":   message,
		"timestamp": internal.GetTimestamp(),
	})
}
