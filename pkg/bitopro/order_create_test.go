package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_CreateOrderLimitBuy(t *testing.T) {
	// satoshi amount
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderLimitBuy(1, "eth_twd", "11", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_CreateOrderLimitSell(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderLimitSell(1, "eth_twd", "1", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_CreateOrderMarketBuy(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderMarketBuy(1, "eth_twd", "1000"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_CreateOrderMarketSell(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CreateOrderMarketSell(1, "eth_twd", "1"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
