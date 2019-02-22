package path

import (
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
}