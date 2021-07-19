package main

import (
	"errors"
	"fmt"
	"testing"
)

func DoMath(num1, num2 int, op string) (int, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 + num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		} else {
			return num1 / num2, nil
		}
	//
	default:
		return 0, fmt.Errorf("unknown operator %s", op)
	}
}

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_divisioin", 2, 0, "/", 0, "division by zero"},
	}

	for _, d := range data {
		t.Run(d.name, func(tt *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				tt.Errorf("Expected %d , got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				tt.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
