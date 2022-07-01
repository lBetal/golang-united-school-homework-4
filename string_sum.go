package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

//
// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input s3ile

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}
	if !isNotNum(input) {
		return "", fmt.Errorf("input have any simbols")
	}
	firtsRegexp, err := regexp.Compile(`(-|\+){0,1}\d+`)
	if err != nil {
		return "", fmt.Errorf("regexp not compile: %w", err)
	}
	arr := firtsRegexp.FindAllString(input, -1)
	if len(arr) != 2 {
		return "", fmt.Errorf("want's two operand: %w", errorNotTwoOperands)
	}
	firstOperand, err := strconv.Atoi(arr[0])
	if err != nil {
		return "", fmt.Errorf("conversion first operand: %w", errorNotTwoOperands)
	}
	secondOperand, err := strconv.Atoi(arr[1])
	if err != nil {
		return "", fmt.Errorf("conversion second operand: %w", errorNotTwoOperands)
	}
	output = strconv.Itoa(firstOperand + secondOperand)
	return output, nil
}

func isNotNum(s string) bool {
	var count int
	for _, val := range s {
		if val == '+' || val == '-' {
			count++
			continue
		}
		if count > 2 {
			return false
		}
		if val < '0' || val > '9' {
			return false
		}
	}
	return true
}
