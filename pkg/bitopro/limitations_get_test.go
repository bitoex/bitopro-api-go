package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetLimitations(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetLimitaionsFees(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(GetPubClient().GetLimitaionsFees(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
