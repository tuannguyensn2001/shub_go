package timepkg

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestParseHour(t *testing.T) {

	t.Run("valid hour", func(t *testing.T) {
		_, err := ParseHour("19:00")

		assert.Equal(t, nil, err)

		_, err = ParseHour("02:00")

		assert.Equal(t, nil, err)

	})

	t.Run("invalid hour", func(t *testing.T) {
		val, err := ParseHour("abc")

		assert.Equal(t, nil, val)
		assert.NotEqual(t, nil, err)

	})
}

func TestParseDate(t *testing.T) {
	t.Run("valid date", func(t *testing.T) {
		_, err := ParseDate("16/05/2022")

		assert.Equal(t, nil, err)

	})

	t.Run("invalid date", func(t *testing.T) {
		val, err := ParseHour("abc")

		assert.Equal(t, nil, val)
		assert.NotEqual(t, nil, err)

	})
}
