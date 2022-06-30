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

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input s3ile

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("Empty input: %w", errorEmptyInput)
	}
	firtsRegexp, err := regexp.Compile(`(-|\+){0,1}\d+`)
	if err != nil {
		return "", fmt.Errorf("Regexp not compile: %w", err)
	}
	arr := firtsRegexp.FindAllString(input, -1)
	if len(arr) != 2 {
		return "", fmt.Errorf("Want's two operand: %w", errorNotTwoOperands)
	}
	firstOperand, err := strconv.Atoi(arr[0])
	if err != nil {
		return "", fmt.Errorf("Error in first operand: %w", err)
	}
	secondOperand, err := strconv.Atoi(arr[1])
	if err != nil {
		return "", fmt.Errorf("Error in second operand: %w", err)
	}
	output = strconv.Itoa(firstOperand + secondOperand)
	return
}
