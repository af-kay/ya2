package ya2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreaSquare(t *testing.T) {
	fn, ok := area(square)

	assert.Equal(t, ok, true)
	assert.Equal(t, fn(4), 16.0)
}

func TestAreaCircle(t *testing.T) {
	fn, ok := area(circle)

	assert.Equal(t, ok, true)
	assert.Equal(t, fn(8), 8*8*3.14)
}

func TestAreaTriangle(t *testing.T) {
	fn, ok := area(triangle)

	assert.Equal(t, ok, true)
	assert.Equal(t, fn(4), 6.0)
}

func TestAreaUnknown(t *testing.T) {
	fn, ok := area(10_000)

	assert.Equal(t, ok, false)
	assert.Panics(t, func() { fn(10_000) })
}
