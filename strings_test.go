package random

import (
	"testing"
)

func TestRandomAlpha(t *testing.T) {
	str1 := RandomAlphaNum(6)
	str2 := RandomAlphaNum(6)
	if str1 == str2 {
		t.Errorf("Random alphanumeric strings are not random [str1=%s, str2=%s]", str1, str2)
	}

}
