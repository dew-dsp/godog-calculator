package calc_test

import (
	"flag"
	"os"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	calc "github.com/dewdsp/godog-calculator"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

type CalcSuite struct {
	*godog.Suite
	calc *calc.Calculator
}

func (cs *CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs *CalcSuite) iPress(arg1 int) error {
	cs.calc.Press(arg1)
	return nil
}

func (cs *CalcSuite) iAdd(arg1 int) error {
	cs.calc.Add(arg1)
	return nil
}

func iSubtract(arg1 int) error {
	return godog.ErrPending
}

func theResultShouldBe(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(suite *godog.Suite) {
	s := CalcSuite{
		calc:  new(calc.Calculator),
		Suite: suite,
	}
	suite.Step(`^calculator is cleared$`, s.calculatorIsCleared())
	suite.Step(`^i press (\d+)$`, s.iPress(2))
	suite.Step(`^i add (\d+)$`, s.iAdd(5))
	suite.Step(`^i subtract (\d+)$`, iSubtract)
	suite.Step(`^the result should be (\d+)$`, theResultShouldBe)
}
