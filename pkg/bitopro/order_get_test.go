package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrder(t *testing.T) {
	var resp []byte
	var err error
	getAuthResp, err := getAuthClient().GetOrder("eth_twd", "384453381")
	if err != nil {
		t.Errorf("getAuthResp err %+v\n", err)
	}
	if resp, err = json.MarshalIndent(getAuthResp, "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("resp=%+v\n", string(resp))
	}
}
