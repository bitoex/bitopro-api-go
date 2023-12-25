package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetTradingPairInfos(t *testing.T) {
	if json, err := json.MarshalIndent(GetPubClient().GetTradingPairInfos(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}

	if json, err := json.MarshalIndent(getAuthClient().GetTradingPairInfos(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
