package bitopro

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestAuthAPI_GetDepositHistory(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetDepositHistory("usdt_twd", url.Values{}), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
