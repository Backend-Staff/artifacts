package functions

import (
	"sort"
	"strings"
)

// Apply applies fn to every item and returns a new slice.
func Apply[T, R any](items []T, fn func(T) R) []R {
	// TODO
	return nil
}

// Filter returns only items where predicate returns true.
func Filter[T any](items []T, predicate func(T) bool) []T {
	// TODO
	return nil
}

// Reduce folds items into one value.
func Reduce[T, R any](items []T, initial R, fn func(R, T) R) R {
	// TODO
	return initial
}

// MakeCounter returns a closure that increments and returns a counter.
func MakeCounter() func() int {
	// TODO
	return func() int { return 0 }
}

// Fibonacci returns the nth Fibonacci number recursively.
func Fibonacci(n int) int {
	// TODO: handle n <= 0
	return 0
}

// SortStrings sorts strings alphabetically, case-insensitively.
func SortStrings(items []string) []string {
	// TODO: do not mutate the input slice
	return nil
}

// Compose returns f(g(x)).
func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	// TODO
	return nil
}

// VariadicMax returns the largest number; return an error for no arguments.
func VariadicMax(nums ...int) (int, error) {
	// TODO
	return 0, nil
}

// CleanWords splits text, removes empty strings, and lowercases words.
func CleanWords(text string) []string {
	// TODO
	return nil
}

var _ = sort.Strings
var _ = strings.ToLower
