package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrder(t *testing.T) {
	var resp []byte
	var err error
	if resp, err = json.MarshalIndent(getAuthClient().GetOrder("eth_twd", 2559364785), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("resp=%+v\n", string(resp))
	}
}
