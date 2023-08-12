package ya2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeferGlobalManip(t *testing.T) {
	expected := Global

	globalDeferManip()

	assert.Equal(t, expected, Global)

	globalDeferManip()

	assert.Equal(t, expected, Global)
}
