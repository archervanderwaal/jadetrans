// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package path

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	notExistPath := "/.jadetrans/notexist"
	if Exists(notExistPath) {
		t.Errorf("not expected")
	}
	existPath := "/Users"
	if !Exists(existPath) {
		t.Errorf("not expected")
	}
	fmt.Println(Exists(notExistPath), Exists(existPath))
}
