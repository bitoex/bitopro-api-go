package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GeTradingPairs(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetTradingpairs(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(GetPubClient().GetTradingpairs(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
