// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package path

import (
	"path/filepath"
	"os"
	"github.com/mitchellh/go-homedir"
)

// Cache the location of the home directory.
var homeDir = ""

// Home returns the Jadetrans home directory.
func Home() string {
	if homeDir != "" {
		return homeDir
	}

	if h, err := homedir.Dir(); err == nil {
		homeDir = filepath.Join(h, ".jadetrans")
	} else {
		cwd, err := os.Getwd()
		if err == nil {
			homeDir = filepath.Join(cwd, ".jadetrans")
		} else {
			homeDir = ".jadetrans"
		}
	}
	return homeDir
}

// SetHome sets the home directory for Jadetrans.
func SetHome(home string) {
	homeDir = home
}

// Exists returns does a file or folder exist
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
