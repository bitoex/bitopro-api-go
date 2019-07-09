package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_CreateOrderLimitBuy(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderLimitBuy("eth_twd", "0.1", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_CreateOrderLimitSell(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderLimitSell("eth_twd", "0.1", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_CreateOrderMarketBuy(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderMarketBuy("eth_twd", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_CreateOrderMarketSell(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderMarketSell("eth_twd", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
