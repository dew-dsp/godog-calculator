package calc

// Calculator struct
type Calculator struct {
	result int
}

// Add function for Calculator
func (c *Calculator) Add(x int) {
	c.result += x
}

// Sub function for Calculator
func (c *Calculator) Sub(x int) {
	c.result -= x
}

// Multiply function for Calculator
func (c *Calculator) Multiply(x int) {
	c.result *= x
}

// Press function for Calculator
func (c *Calculator) Press(x int) {
	c.result = x
}

// Clear function for Calculator
func (c *Calculator) Clear() {
	c.result = 0
}

// Result function for Calculator
func (c *Calculator) Result() int {
	return c.result
}
