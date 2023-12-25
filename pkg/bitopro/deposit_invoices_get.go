package bitopro

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/bitoex/bitopro-api-go/internal"
)

type TransactionInfo struct {
	Serial    string `json:"serial"`
	Timestamp string `json:"timestamp"`
	Address   string `json:"address"`
	Amount    string `json:"amount"`
	Fee       string `json:"fee"`
	Total     string `json:"total"`
	Status    string `json:"status"`
	Txid      string `json:"txid"`
	Protocol  string `json:"protocol"`
	ID        string `json:"id"`
}

// DepositHistory struct
type DepositHistory struct {
	Data []TransactionInfo `json:"data,omitempty"`
	StatusCode
}

// GetDepositHistory
func (api *AuthAPI) GetDepositHistory(currency string, queryParams url.Values) *DepositHistory {
	var data DepositHistory

	queryStr := queryParams.Encode()
	endpoint := fmt.Sprintf("%s/%s", "v3/wallet/depositHistory", currency)
	if queryStr != "" {
		endpoint = fmt.Sprintf("%s/%s?%s", "v3/wallet/depositHistory", currency, queryStr)
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
