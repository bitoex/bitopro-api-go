package bitopro

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestAuthAPI_GetWithdrawHistory(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetWithdrawHistory("usdt_twd", url.Values{}), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

// GetWithdrawBySerial, GetWithdrawById
