package bitopro

import (
	"encoding/json"

	"github.com/bitoex/bitopro-api-go/internal"
)

// CurrencyInfo struct
type CurrencyInfo struct {
	Currency            string `json:"currency"`
	WithdrawFee         string `json:"withdrawFee"`
	MinWithdraw         string `json:"minWithdraw"`
	MaxWithdraw         string `json:"maxWithdraw"`
	MaxDailyWithdraw    string `json:"maxDailyWithdraw"`
	Withdraw            bool   `json:"withdraw"`
	Deposit             bool   `json:"deposit"`
	DepositConfirmation string `json:"depositConfirmation"`
}

// Currency struct
type Currency struct {
	Data []CurrencyInfo `json:"data,omitempty"`
	StatusCode
}

func getCurrency(proxy string) *Currency {
	var data Currency

	code, res, err := internal.ReqPublic("v3/provisioning/currencies", proxy)
	if err != nil {
		data.Error = err.Error()
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

func (p *PubAPI) GetCurrency() *Currency {
	return getCurrency(p.proxy)
}

func (a *AuthAPI) GetCurrency() *Currency {
	return getCurrency(a.proxy)
}
