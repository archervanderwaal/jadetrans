// Package utils Provides some tools about time.
package utils

import (
	"strconv"
	"time"
)

// UTCTimestamp returns utc timestamp.
func UTCTimestamp() string {
	current := time.Now()
	timestamp := int(current.UnixNano() / 1000000000)
	return strconv.Itoa(timestamp)
}
