package color_test

import (
	"testing"

	"github.com/derailed/k9s/internal/color"
	"github.com/stretchr/testify/assert"
)

func TestColorize(t *testing.T) {
	uu := map[string]struct {
		s string
		c color.Paint
		e string
	}{
		"white":   {"blee", color.LightGray, "\x1b[37mblee\x1b[0m"},
		"black":   {"blee", color.Black, "\x1b[30mblee\x1b[0m"},
		"default": {"blee", 0, "blee"},
	}

	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, color.Colorize(u.s, u.c))
		})
	}
}
