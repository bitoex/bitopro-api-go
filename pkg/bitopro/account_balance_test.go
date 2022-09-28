package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetAccountBalance(t *testing.T) {
	// internal.SetEndpoint("xxx")

	if json, err := json.MarshalIndent(getAuthClient().GetAccountBalance(), "", ""); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
