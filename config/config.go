// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package config

import (
	"fmt"
	"github.com/archervanderwaal/jadetrans/path"
	"github.com/aybabtme/rgbterm"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const (
	configFile = "jadetrans.yaml"
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
		if err != nil {
			log.Println(rgbterm.FgString("Internal error "+err.Error(), 255, 0, 0))
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprintf(file, "youdao:\n  appKey: %s\n  appSecret: %s\n", "", "")
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
	AppKey    string `yaml:"appKey"`
	AppSecret string `yaml:"appSecret"`
}

type Config struct {
	Youdao YoudaoConfig `yaml:"youdao"`
}

// LoadConfig returns jadetrans config.
func LoadConfig() *Config {
	var settings Config
	configFile, err := ioutil.ReadFile(location())
	if err != nil {
		log.Println(rgbterm.FgString("Internal error "+err.Error(), 255, 0, 0))
		os.Exit(1)
	}
	yaml.Unmarshal(configFile, &settings)
	return &settings
}
