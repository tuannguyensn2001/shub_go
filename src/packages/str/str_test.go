package strpkg

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRandom(t *testing.T) {
	result := Random(5)
	assert.Equal(t, 5, len(result))
}
