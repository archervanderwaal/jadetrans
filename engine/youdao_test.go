// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package engine

import (
	"github.com/archervanderwaal/jadetrans/config"
	"testing"
)

var engine *YoudaoEngine

var conf *config.Config

func init() {
	conf = config.LoadConfig()
}

func TestNewYoudaoEngine(t *testing.T) {
	engine, _ = NewYoudaoEngine(
		"peace", "auto", "auto",
		"0", conf)
}

func TestYoudaoEngine_Query(t *testing.T) {
	if engine == nil {
		engine, _ = NewYoudaoEngine(
			"peace", "auto", "auto",
			"0", conf)
	}
	engine.Query()
}
