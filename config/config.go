// Package config Provides loading local configuration.
//
// Source code and other details for the project are available at GitHub:
//
// https://github.com/archervanderwaal/jadetrans
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"

	"github.com/archervanderwaal/jadetrans/path"
	"github.com/aybabtme/rgbterm"
)

const (
	configFile = "jadetrans.yaml"
)

// Setup creates the home dir and config file.
func setup() {
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
}

// Location returns the location of the config file.
func location() string {
	configFile := filepath.Join(path.Home(), configFile)
	setup()
	return configFile
}

// YoudaoConfig configuration information encapsulated with youdao Translation.
type YoudaoConfig struct {
	AppKey    string `yaml:"appKey"`
	AppSecret string `yaml:"appSecret"`
}

// Config jadetrans global config.
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
