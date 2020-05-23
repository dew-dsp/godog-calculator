package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Add(t *testing.T) {
	c := new(Calculator)
	c.Press(10)
	c.Add(10)
	assert.Equal(t, 20, c.Result())
}

func TestCalculator_Sub(t *testing.T) {
	c := new(Calculator)
	c.Press(30)
	c.Sub(10)
	assert.Equal(t, 20, c.Result())
}

func TestCalculator_Multiply(t *testing.T) {
	c := new(Calculator)
	c.Press(2)
	c.Multiply(4)
	assert.Equal(t, 8, c.Result())
}

func TestCalculator_Press(t *testing.T) {
	c := new(Calculator)
	c.Press(3)
	assert.Equal(t, 3, c.Result())
}

func TestCalculator_Clear(t *testing.T) {
	c := new(Calculator)
	c.Clear()
	assert.Equal(t, 0, c.Result())
}
