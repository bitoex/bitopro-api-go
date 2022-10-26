package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrder(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOrder("eth_twd", 1), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("resp=%+v\n", string(json))
	}
}
