// Package path provides some tool functions related to the file system.
//
// Source code and other details for the project are available at GitHub:
//
// https://github.com/archervanderwaal/jadetrans
package path

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// Cache the location of the home directory.
var homeDir = ""

// The directory of the project configuration file.
const folder = ".jadetrans"

// Home returns the Jadetrans home directory.
func Home() string {
	if homeDir != "" {
		return homeDir
	}

	if h, err := homedir.Dir(); err == nil {
		homeDir = filepath.Join(h, folder)
	} else {
		cwd, err := os.Getwd()
		if err == nil {
			homeDir = filepath.Join(cwd, folder)
		} else {
			homeDir = folder
		}
	}
	return homeDir
}

// Exists returns does a file or folder exist.
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
