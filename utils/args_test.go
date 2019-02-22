// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package utils

import (
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	os.Args = []string{
		"",
		"-help",
		"-version",
		"I",
		"Love",
		"You",
		"-egine=youdao",
	}
	words, args := ParseArgs(os.Args)
	if len(words) == 3 && len(args) == 3 {
		return
	}
	t.Fail()
}
