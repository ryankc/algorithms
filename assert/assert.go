package assert

import "testing"

func Equal(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Equality Check failed: Expected: %v, Actual: %v", expected, actual)
	}
}
