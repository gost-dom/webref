package elements_test

import (
	"testing"

	"github.com/gost-dom/webref/elements"
	"github.com/stretchr/testify/assert"
)

func TestGetTagNameForInterface(t *testing.T) {
	assert := assert.New(t)
	e, err := elements.Load("html")
	assert.NoError(err)
	actual, ok := e.GetTagNameForInterface("HTMLAnchorElement")
	assert.True(ok)
	assert.Equal("a", actual)
}

func TestGetInterfaceForTagName(t *testing.T) {
	assert := assert.New(t)
	e, err := elements.Load("html")
	assert.NoError(err)
	actual, ok := e.GetTagNameForInterface("HTMLAnchorElement")
	assert.True(ok)
	assert.Equal("a", actual)
}
