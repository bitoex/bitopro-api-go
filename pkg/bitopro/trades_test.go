package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetTrades(t *testing.T) {
	if json, err := json.MarshalIndent(GetPubClient().GetTrades("eth_twd"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(getAuthClient().GetTrades("eth_twd"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
