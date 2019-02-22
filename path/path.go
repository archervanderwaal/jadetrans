// Package path Path-dependent functions.
package path

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
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

// Exists returns does a file or folder exist
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
