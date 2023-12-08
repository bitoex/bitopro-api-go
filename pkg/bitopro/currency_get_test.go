package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetCurrenvy(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetCurrency(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(GetPubClient().GetCurrency(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
