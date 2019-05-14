package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetTicker(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetTicker("eth_twd"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestPubAPI_GetTicker(t *testing.T) {
	if json, err := json.MarshalIndent(GetPubClient().GetTicker("eth_twd"), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
