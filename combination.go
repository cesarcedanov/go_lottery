package main

import (
	"errors"
	"sort"
	"strings"
)

// Combination is a array of number
type Combination struct {
	Numbers []string
}

// NewCombination extract a combination of values from an Array of string
func NewCombination(numbers []string, length int) (Combination, error) {
	if length < 1 {
		return Combination{}, errors.New("Invalid combination length")
	}

	return Combination{Numbers: numbers[:length]}, nil
}

// toString concat the all value from the string array into a single string
func (c Combination) ToString() string {
	return strings.Join(c.Numbers, ", ")
}

// sort will sort the value from the string array
func (c Combination) Sort() {
	sort.Strings(c.Numbers)
}