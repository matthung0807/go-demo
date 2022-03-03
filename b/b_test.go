package b

import "testing"

func TestB(t *testing.T) {
	result := B(1)
	if result != 1 {
		t.Errorf("expect %d, but %d", 1, result)
	}
}
