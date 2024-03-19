package postfix

import (
	"errors"
	"strconv"
	"strings"
)

// EvalPostfix обчислює значення постфіксного виразу.
func EvalPostfix(expr string) (float64, error) {
	tokens := strings.Fields(expr)
	stack := []float64{}

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression: insufficient operands")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0, errors.New("division by zero")
				}
				result = operand1 / operand2
			case "^":
				result = pow(operand1, operand2)
			}
			stack = append(stack, result)
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression: too many operands")
	}

	return stack[0], nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func pow(base, exponent float64) float64 {
	result := 1.0
	for i := 0; i < int(exponent); i++ {
		result *= base
	}
	return result
}
