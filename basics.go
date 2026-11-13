package basics

import (
	"errors"
	"fmt"
	"strings"
)


func Sum(nums ...int) int {
	
	total := 0
	for _, num := range nums{
		total += num
	}
	return total
}

// Reverse returns a reversed copy of s.
func Reverse(s string) string {
	// TODO: handle Unicode correctly using runes
	// 
	// 
	return s
}

// WordCount returns how many times each whitespace-separated word appears.
func WordCount(text string) map[string]int {
	// TODO: use strings.Fields and normalize case
	return nil
}

// FizzBuzz returns:
// "Fizz" for multiples of 3
// "Buzz" for multiples of 5
// "FizzBuzz" for multiples of both
// otherwise the number as a string.
func FizzBuzz(n int) string {
	// TODO
	return ""
}

// Unique removes duplicate values while preserving first-seen order.
func Unique[T comparable](items []T) []T {
	// TODO
	return nil
}

// ParsePositive parses a positive integer.
// Return an error for empty input, invalid numbers, zero, or negatives.
func ParsePositive(s string) (int, error) {
	// TODO: use strconv.Atoi and errors.New / fmt.Errorf
	return 0, errors.New("not implemented")
}

// FormatNames turns ["ada", "grace"] into "Ada, Grace".
func FormatNames(names []string) string {
	// TODO: title-case each name and join with ", "
	return ""
}

// Describe demonstrates fmt verbs.
func Describe(name string, age int, score float64) string {
	// TODO: return exactly: "name=Ada age=30 score=97.50"
	return fmt.Sprintf("", name, age, score)
}

// SafeDivide divides a by b.
// Return an error when b is zero.
func SafeDivide(a, b float64) (float64, error) {
	// TODO
	return 0, nil
}

// NormalizeEmail trims spaces and lowercases an email.
func NormalizeEmail(email string) string {
	// TODO
	return strings.TrimSpace(email)
}
