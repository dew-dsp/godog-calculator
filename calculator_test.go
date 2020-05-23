package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IndexTests struct {
	start    int
	x        int
	expected int
}

func TestCalculator_Add(t *testing.T) {
	var addResults = []IndexTests{
		{1, 2, 3},
		{5, 3, 8},
		{0, 2, 2},
	}

	for _, test := range addResults {
		c := new(Calculator)
		c.Press(test.start)
		c.Add(test.x)
		assert.Equal(t, test.expected, c.Result())
	}
}

func TestCalculator_Sub(t *testing.T) {
	var subResults = []IndexTests{
		{20, 5, 15},
		{-2, -3, 1},
	}

	for _, test := range subResults {
		c := new(Calculator)
		c.Press(test.start)
		c.Sub(test.x)
		assert.Equal(t, test.expected, c.Result())
	}
}

func TestCalculator_Multiply(t *testing.T) {
	var multiplyResults = []IndexTests{
		{2, 5, 10},
		{10, -8, -80},
		{-2, -4, 8},
	}

	for _, test := range multiplyResults {
		c := new(Calculator)
		c.Press(test.start)
		c.Multiply(test.x)
		assert.Equal(t, test.expected, c.Result())
	}
}

func TestCalculator_Clear(t *testing.T) {
	c := new(Calculator)
	c.Clear()
	assert.Equal(t, 0, c.Result())
}

func TestCalculator_Press(t *testing.T) {
	type pressTests struct {
		x        int
		expected int
	}

	var pressResults = []pressTests{
		{3, 3},
		{4, 4},
		{0, 0},
	}
	c := new(Calculator)
	for _, test := range pressResults {
		c.Press(test.x)
		assert.Equal(t, test.expected, c.Result())
	}
}

func TestBla(t *testing.T) {
	c := new(Calculator)
	c.Press(5)
	c.Press(3)
	c.Add(4)

	assert.Equal(t, 7, c.Result())
}
