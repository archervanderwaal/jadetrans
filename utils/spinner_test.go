package utils

import (
	"testing"
	"time"
)

var ids [3]uint64

func TestNewDefaultSpinnerAndStart(t *testing.T) {
	for i := 0; i < 3; i++ {
		ids[i] = NewDefaultSpinnerAndStart("Querying...")
	}
}

func TestStopSpinner(t *testing.T) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		StopSpinner(ids[i])
	}
}
