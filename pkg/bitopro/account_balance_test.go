package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetAccountBalance(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetAccountBalance(), "", ""); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
