package cast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	assert.Equal(t, 123, ToInt("123", WithBase(10)))
}
