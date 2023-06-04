package internal

import (
	"encoding/base64"
	"encoding/json"
)

func getNonPostPayload(identity string, nonce int64) string {
	payload, _ := json.Marshal(map[string]interface{}{
		"identity": identity,
		"nonce":    nonce,
	})

	return base64.StdEncoding.EncodeToString(payload)
}

func getPostPayload(body map[string]interface{}) (string, string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return "", "", err
	}

	return string(payload), base64.StdEncoding.EncodeToString(payload), nil
}
