// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
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
