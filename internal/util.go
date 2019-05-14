package internal

import (
	"time"
)

// GetTimestamp func
func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
