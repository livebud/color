package color_test

import (
	"testing"

	"github.com/livebud/color"
	"github.com/matryer/is"
)

func TestColor(t *testing.T) {
	is := is.New(t)
	color := color.New()
	is.Equal(color.Blue("hello"), "\x1b[34mhello\x1b[0m")
	is.Equal(color.Red("hello"), "\x1b[31mhello\x1b[0m")
}

func TestIgnore(t *testing.T) {
	is := is.New(t)
	color := color.Ignore()
	is.Equal(color.Blue("hello"), "hello")
	is.Equal(color.Red("hello"), "hello")
}
