package internal

import (
	"testing"
)

func Test_getNonPostPayload(t *testing.T) {
	if getNonPostPayload("support@bitoex.com", 1554380909131) != "eyJpZGVudGl0eSI6InN1cHBvcnRAYml0b2V4LmNvbSIsIm5vbmNlIjoxNTU0MzgwOTA5MTMxfQ==" {
		t.Error("not match")
	}
}

func Test_getPostPayload(t *testing.T) {
	if _, payload := getPostPayload(map[string]interface{}{
		"action":    "buy",
		"type":      "limit",
		"price":     "1.123456789",
		"amount":    "666",
		"timestamp": 1554380909131,
	}); payload != "eyJhY3Rpb24iOiJidXkiLCJhbW91bnQiOiI2NjYiLCJwcmljZSI6IjEuMTIzNDU2Nzg5IiwidGltZXN0YW1wIjoxNTU0MzgwOTA5MTMxLCJ0eXBlIjoibGltaXQifQ==" {
		t.Error("not match")
	}
}
