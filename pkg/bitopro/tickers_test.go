package bitopro

import (
	"encoding/json"
	"testing"
)

func TestPubAPI_GetTickers(t *testing.T) {
	if json, err := json.MarshalIndent(GetPubClient().GetTickers(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}

func TestAuthAPI_GetTickers(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetTickers(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
