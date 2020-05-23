package calc_test

import (
	calc "github.com/sitthakarn/godog-tutorial"
)

type CalcSuite struct{
	*godog.Suite
	calc *calc.Calculator
}

func (cs *CalcSuite) calculatorIsCleared() error {
	cs
	return godog.ErrPending
}

func iPress(arg1 int) error {
	return godog.ErrPending
}

func iAdd(arg1 int) error {
	return godog.ErrPending
}

func iSubtract(arg1 int) error {
	return godog.ErrPending
}

func theResultShouldBe(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^calculator is cleared$`, calculatorIsCleared)
	s.Step(`^i press (\d+)$`, iPress)
	s.Step(`^i add (\d+)$`, iAdd)
	s.Step(`^i subtract (\d+)$`, iSubtract)
	s.Step(`^the result should be (\d+)$`, theResultShouldBe)
}
