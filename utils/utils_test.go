package utils_test

import (
	"testing"

	u "github.com/MeztliRA/gemdot/utils"
)

const (
	expectedStr = "gemdot"
	str         = expectedStr + "	 \n"
)

func TestTrimString(t *testing.T) {
	got := u.TrimString(str)
	if got != expectedStr {
		t.Errorf("failed to parse string")
	}
}
