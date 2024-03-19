package postfix

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type PostfixSuite struct{}

var _ = Suite(&PostfixSuite{})

func (s *PostfixSuite) TestEvalPostfix(c *C) {
	cases := []struct {
		expr     string
		expected float64
		err      bool
	}{
		{"4 2 - 3 * 5 +", 11, false},
		{"10 2 ^", 100, false},
		{"5 0 /", 0, true},         // Division by zero
		{"5 *", 0, true},           // Invalid expression
		{"10 5 -", 5, false},       // 10 - 5 = 5
		{"20 10 - 5 *", 50, false}, // (20 - 10) * 5 = 50
		{"15 -5 +", 10, false},     // 15 + (-5) = 10
	}

	for _, tc := range cases {
		result, err := EvalPostfix(tc.expr)
		if tc.err {
			c.Assert(err, NotNil)
		} else {
			c.Assert(err, IsNil)
			c.Assert(result, Equals, tc.expected)
		}
	}
}

// ExampleEvalPostfix illustrates the use of EvalPostfix in documentation.
func ExampleEvalPostfix() {
	result, _ := EvalPostfix("4 2 - 3 * 5 +")
	_ = result // Ігноруємо результат
	// Виведення результата:
	fmt.Println(result)
	// Output: 11
}
