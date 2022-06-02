package models

import (
	"github.com/magiconair/properties/assert"
	"shub_go/src/enums"
	"testing"
)

func TestGetIsShow(t *testing.T) {
	t.Run("get true", func(t *testing.T) {
		post := Post{}
		post.SetIsShow(enums.IsShow)

		assert.Equal(t, true, post.GetIsShow())
	})

	t.Run("get false", func(t *testing.T) {
		post := Post{}
		post.SetIsShow(enums.NotShow)

		assert.Equal(t, false, post.GetIsShow())
	})
}