// Package engine implements translation support for the translation engine
// so that the need for translating text can be achieved by creating specific
// translation engine.
//
// Source code and other details for the project are available at GitHub:
//
// https://github.com/archervanderwaal/jadetrans
package engine

// Engine Represents a translation engine.
type Engine interface {
	Query() (res string, err error)
}
