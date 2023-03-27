package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"udmey/module06/internal/util"
)

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)

	isPalindromeTestOk := func(input int, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			token := util.IsPalindrome(input)
			a.Equal(true, token)
		}
	}

	isPalindromeTestFail := func(input int, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			token := util.IsPalindrome(input)
			a.Equal(false, token)
		}
	}

	t.Run("even palindrome value", isPalindromeTestOk(1221, true))
	t.Run("odd palindrome value", isPalindromeTestOk(12321, true))
	t.Run("negative value", isPalindromeTestOk(-1111, true))

	t.Run("random numbers value", isPalindromeTestFail(1234, false))

}
