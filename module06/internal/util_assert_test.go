package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPosInt(t *testing.T) {
	token, err := ReverseNumber(1212)
	assert.NoError(t, err)
	assert.Equal(t, token, 2121, "should be equal")
}

func TestNegInt(t *testing.T) {
	token, err := ReverseNumber(-1212)
	assert.NoError(t, err)
	assert.Equal(t, token, -2121, "should be equal")
}

func TestStr(t *testing.T) {
	token, err := ReverseNumber("1212")
	assert.NoError(t, err)
	assert.Equal(t, token, "2121", "should be equal")
}

func TestInvalType(t *testing.T) {
	token, err := ReverseNumber(12.22)

	assert.ErrorContains(t, err, "incorrect type")
	assert.Nil(t, token)
}
