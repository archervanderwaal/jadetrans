// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package utils

import "github.com/satori/go.uuid"

func Uuid() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}
