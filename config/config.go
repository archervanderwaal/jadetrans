// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package config

import (
	"sync"
	"github.com/archervanderwaal/jadetrans/path"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v1"
)

const (
	configFile	=    "jadetrans.yaml"
)

var isSetup bool

var setupMutex sync.Mutex

// Setup creates the home dir and config file.
func setup() {
	setupMutex.Lock()
	defer setupMutex.Unlock()

	if isSetup {
		return
	}
	if !path.Exists(path.Home()) {
		os.Mkdir(path.Home(), 0755)
	}
	if !path.Exists(filepath.Join(path.Home(), configFile)) {
		file, err := os.Create(filepath.Join(path.Home(), configFile))
		defer file.Close()
		if err == nil {
			fmt.Fprintf(file, "youdao:\n  appId: %s\n  appKey: %s\n", "", "")
		}
	}
	isSetup = true
}

// SetupReset resets if setup has been completed. The next time setup is run
// it will attempt a full setup.
func setupReset() {
	isSetup = false
}

// Location returns the location of the config file.
func location() string {
	configFile := filepath.Join(path.Home(), configFile)
	setup()
	return configFile
}

type YoudaoConfig struct {
	AppId    	string `yaml:"appId"`
	AppKey   	string `yaml:"appKey"`
}

type Config struct {
	Youdao 		YoudaoConfig `yaml:"youdao"`
}

func LoadConfig() *Config {
	var settings Config
	configFile, err := ioutil.ReadFile(location())
	if err == nil {
		yaml.Unmarshal(configFile, &settings)
	}
	return &settings
}

