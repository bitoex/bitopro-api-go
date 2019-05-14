package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrderBook(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOrderBook("eth_twd"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestPubAPI_GetOrderBook(t *testing.T) {
	if json, err := json.MarshalIndent(GetPubClient().GetOrderBook("eth_twd"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
