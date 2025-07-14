package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWithAssert(t *testing.T) {
	got := Add(2, 3)
	want := 5
	assert.Equal(t, want, got, "they should be equal")

	// Example of another assertion
	assert.NotEqual(t, 6, got, "they should not be equal")
}
