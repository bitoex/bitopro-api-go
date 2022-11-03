package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OrderInfo struct
type OrderInfo struct {
	ID                string `json:"id"`
	Pair              string `json:"pair"`
	Price             string `json:"price"`
	AvgExecutionPrice string `json:"avgExecutionPrice,omitempty"`
	Action            string `json:"action,omitempty"`
	Type              string `json:"type,omitempty"`
	Timestamp         int64  `json:"timestamp,omitempty"`
	Status            int    `json:"status,omitempty"`
	CreatedTimestamp  int64  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  int64  `json:"updatedTimestamp,omitempty"`
	OriginalAmount    string `json:"originalAmount,omitempty"`
	RemainingAmount   string `json:"remainingAmount,omitempty"`
	ExecutedAmount    string `json:"executedAmount,omitempty"`
	Fee               string `json:"fee,omitempty"`
	FeeSymbol         string `json:"feeSymbol,omitempty"`
	BitoFee           string `json:"bitoFee,omitempty"`
	Total             string `json:"total,omitempty"`
	Seq               string `json:"seq,omitempty"`
	TimeInForce       string `json:"timeInForce,omitempty"`
	Error             error
	StatusCode
}

// GetOrder Ref. https://developer.bitopro.com/docs#operation/getOrderStatus
func (api *AuthAPI) GetOrder(pair string, orderID int) *OrderInfo {
	var data OrderInfo

	code, res := internal.ReqWithoutBody(api.identity, api.Key, api.secret, "GET", fmt.Sprintf("%s/%s/%d", "v3/orders", pair, orderID))
	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = err
	}

	data.Code = code

	return &data
}
