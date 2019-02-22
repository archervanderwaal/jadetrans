// Package utils Provides some tools about uuid.
package utils

import "github.com/satori/go.uuid"

// UUID returns uuid.
func UUID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}
