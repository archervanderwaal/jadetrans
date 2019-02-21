// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package config

import (
	"testing"
	"fmt"
)

func TestLoadConfig(t *testing.T) {
	config := LoadConfig()
	fmt.Println(config.Youdao.AppId, config.Youdao.AppKey)
}