package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return input, fmt.Errorf("got error: %w", errorEmptyInput)
	}

	var element string
	var elements []string

	// remove spaces from input string
	input = strings.Join(strings.Fields(input), "")

	// parsing the string into array with elements with initial signs
	for i, v := range input {
		// If the character + or - is not on the first position, then add the item to the array
		if (v == '+' || v == '-') && i > 0 {
			elements = append(elements, element)
			element = ""
		}

		element = element + string(v)

		// if the character is the last one, add the item to the array
		if i == len(input)-1 {
			elements = append(elements, element)
		}
	}

	if len(elements) < 2 || len(elements) > 2 {
		return output, fmt.Errorf("%s: %w", input, errorNotTwoOperands)
	}

	out := 0
	for _, v := range elements {
		s, err := strconv.Atoi(v)
		if err != nil {
			return output, fmt.Errorf("%s: %w", input, err)
		}
		out += s
	}
	return strconv.Itoa(out), nil
}
