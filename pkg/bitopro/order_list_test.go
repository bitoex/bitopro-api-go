package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrderList(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOrderList("eth_twd", false, 1), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(getAuthClient().GetOrderList("eth_twd", true, 1), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
