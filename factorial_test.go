package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {

	t.Run("factorial de 0 debe dar 1", func(t *testing.T) {
		value := 0
		expected := 1
		resultado := factorial(value)
		assert.Equal(t, expected, resultado)
	})

	t.Run("factorial de 1 debe dar 1", func(t *testing.T) {
		value := 1
		expected := 1
		resultado := factorial(value)
		assert.Equal(t, expected, resultado)
	})

	t.Run("factorial de 2 debe dar 2", func(t *testing.T) {
		value := 2
		expected := 2
		resultado := factorial(value)
		assert.Equal(t, expected, resultado)
	})

	t.Run("factorial de 3 debe dar 6", func(t *testing.T) {
		value := 3
		expected := 6
		resultado := factorial(value)
		assert.Equal(t, expected, resultado)
	})

	t.Run("factorial de 3 debe dar 6", func(t *testing.T) {
		value := 4
		expected := 24
		resultado := factorial(value)
		assert.Equal(t, expected, resultado)
	})
}
