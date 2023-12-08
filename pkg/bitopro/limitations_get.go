package bitopro

import (
	"encoding/json"

	"github.com/bitoex/bitopro-api-go/internal"
)

type CryptocurrencyDepositFeeAndConfirmationInfo struct {
	BlockchainConfirmationRequired string `json:"blockchainConfirmationRequired"`
	Currency                       string `json:"currency"`
	GeneralDepositFees             string `json:"generalDepositFees"`
}

type OrderFeesAndLimitationsInfo struct {
	Base                          string `json:"base"`
	BasePriceOfDigits             string `json:"basePriceOfDigits"`
	Maintain                      bool   `json:"maintain"`
	MaximumOrderAmount            string `json:"maximumOrderAmount"`
	MinimumOrderAmount            string `json:"minimumOrderAmount"`
	MinimumOrderAmountBase        string `json:"minimumOrderAmountBase"`
	MinimumOrderLimitQuoteAmount  string `json:"minimumOrderLimitQuoteAmount"`
	MinimumOrderMarketQuoteAmount string `json:"minimumOrderMarketQuoteAmount"`
	MinimumOrderNumberOfDigits    string `json:"minimumOrderNumberOfDigits"`
	OrderAmountOfDigits           string `json:"orderAmountOfDigits"`
	OrderBookQuotePrecision       string `json:"orderBookQuotePrecision"`
	OrderBookQuoteScaleLevel      string `json:"orderBookQuoteScaleLevel"`
	OrderOpenLimit                string `json:"orderOpenLimit"`
	Pair                          string `json:"pair"`
	Quote                         string `json:"quote"`
	QuotePriceOfDigits            string `json:"quotePriceOfDigits"`
}

type TradingFeeRateInfo struct {
	Rank             int    `json:"rank"`
	TwdVolumeSymbol  string `json:"twdVolumeSymbol"`
	TwdVolume        string `json:"twdVolume"`
	BitoAmountSymbol string `json:"bitoAmountSymbol"`
	BitoAmount       string `json:"bitoAmount"`
	MakerFee         string `json:"makerFee"`
	TakerFee         string `json:"takerFee"`
	MakerBitoFee     string `json:"makerBitoFee"`
	TakerBitoFee     string `json:"takerBitoFee"`
	RankCondition    string `json:"rankCondition"`
	GridBotMakerFee  string `json:"gridBotMakerFee"`
	GridBotTakerFee  string `json:"gridBotTakerFee"`
}

type RestrictionsOfWithdrawalFeesInfo struct {
	Currency                 string `json:"currency"`
	Fee                      string `json:"fee"`
	MinimumTradingAmount     string `json:"minimumTradingAmount"`
	MaximumTradingAmount     string `json:"maximumTradingAmount"`
	DailyCumulativeMaxAmount string `json:"dailyCumulativeMaximumAmount"`
	Remarks                  string `json:"remarks"`
	Protocol                 string `json:"protocol"`
}

type TTCheckFeesAndLimitationsLevel1Info struct {
	Currency                             string `json:"currency"`
	GenerateDailyCumulativeMaximumAmount string `json:"generateDailyCumulativeMaximumAmount"`
	GenerateMaximumTradingAmount         string `json:"generateMaximumTradingAmount"`
	GenerateMinimumTradingAmount         string `json:"generateMinimumTradingAmount"`
	RedeemDailyCumulativeMaximumAmount   string `json:"redeemDailyCumulativeMaximumAmount"`
}

type TTCheckFeesAndLimitationsLevel2Info struct {
	Currency                             string `json:"currency"`
	GenerateDailyCumulativeMaximumAmount string `json:"generateDailyCumulativeMaximumAmount"`
	GenerateMaximumTradingAmount         string `json:"generateMaximumTradingAmount"`
	GenerateMinimumTradingAmount         string `json:"generateMinimumTradingAmount"`
	RedeemDailyCumulativeMaximumAmount   string `json:"redeemDailyCumulativeMaximumAmount"`
}

// LimitaionsFeesInfo struct
type LimitaionsFeesInfo struct {
	TradingFeeRate                          []TradingFeeRateInfo                          `json:"tradingFeeRate"`
	RestrictionsOfWithdrawalFees            []RestrictionsOfWithdrawalFeesInfo            `json:"restrictionsOfWithdrawalFees"`
	CryptocurrencyDepositFeeAndConfirmation []CryptocurrencyDepositFeeAndConfirmationInfo `json:"cryptocurrencyDepositFeeAndConfirmation"`
	OrderFeesAndLimitations                 []OrderFeesAndLimitationsInfo                 `json:"orderFeesAndLimitations"`
	TTCheckFeesAndLimitationsLevel1         []TTCheckFeesAndLimitationsLevel1Info         `json:"ttCheckFeesAndLimitationsLevel1"`
	TTCheckFeesAndLimitationsLevel2         []TTCheckFeesAndLimitationsLevel2Info         `json:"ttCheckFeesAndLimitationsLevel2"`
	StatusCode
}

func getLimitaionsFees(proxy string) *LimitaionsFeesInfo {
	var data LimitaionsFeesInfo

	code, res, err := internal.ReqPublic("v3/provisioning/limitations-and-fees", proxy)
	if err != nil {
		data.Error = err.Error()
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

func (p *PubAPI) GetLimitaionsFees() *LimitaionsFeesInfo {
	return getLimitaionsFees(p.proxy)
}

func (a *AuthAPI) GetLimitaionsFees() *LimitaionsFeesInfo {
	return getLimitaionsFees(a.proxy)
}
