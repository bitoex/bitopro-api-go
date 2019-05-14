package internal

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

// GetSig func
func getSig(secret, payload string) string {
	h := hmac.New(sha512.New384, []byte(secret))
	h.Write([]byte(payload))

	return hex.EncodeToString(h.Sum(nil))
}
