// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package utils

import (
	"github.com/briandowns/spinner"
	"sync/atomic"
	"time"
)

var (
	// Save the created spinner.
	spinners map[uint64]*spinner.Spinner
	count    uint64 = 0
)

// NewDefaultSpinnerAndStart are defined for create a spinner use default configuration,
// and start it.
func NewDefaultSpinnerAndStart(prefix string) uint64 {
	if spinners == nil {
		spinners = make(map[uint64]*spinner.Spinner)
	}
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Prefix = prefix
	s.Color("green")
	s.Start()
	atomic.AddUint64(&count, 1)
	spinners[count] = s
	return count
}

// StopSpinner are defined for stop spinner according to id.
func StopSpinner(id uint64) {
	if s, ok := spinners[id]; ok {
		s.Stop()
		delete(spinners, id)
	}
}
