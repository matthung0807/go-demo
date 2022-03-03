package a

import "testing"

func TestA_1(t *testing.T) {
	result := A(1)
	if result != 1 {
		t.Errorf("expect %d, but %d", 1, result)
	}
}

func TestA_2(t *testing.T) {
	result := A(-1)
	if result != -1 {
		t.Errorf("expect %d, but %d", -1, result)
	}
}
