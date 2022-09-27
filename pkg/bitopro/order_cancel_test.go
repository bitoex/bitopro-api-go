package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_CancelOrder(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().CancelOrder("eth_twd", 5301507365), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
