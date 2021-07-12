package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	expected := 3
	actial := Plus(2, 1)

	// use Testify assert
	assert.True(t, expected == actial)
	assert.Equal(t, expected, actial)
}
