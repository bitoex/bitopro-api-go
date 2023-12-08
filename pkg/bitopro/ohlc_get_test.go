package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOhlc(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOhlc("usdt_twd", "1d", 1650822974, 1670822974), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(GetPubClient().GetOhlc("usdt_twd", "1d", 1650822974, 1670822974), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
