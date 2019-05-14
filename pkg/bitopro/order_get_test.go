package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrder(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOrder("bito_eth", 2640904509), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
