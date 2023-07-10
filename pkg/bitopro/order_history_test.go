package bitopro

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestAuthAPI_GetOrderHistory(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOrderHistory("btc_usdt", url.Values{}), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
