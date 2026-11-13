package control

import "fmt"

// Grade returns "A", "B", "C", "D", or "F".
func Grade(score int) string {
	// TODO: use if/else if or switch
	return ""
}

// DayType returns "weekend" for Saturday/Sunday, otherwise "weekday".
// Use a switch statement.
func DayType(day string) string {
	// TODO
	return ""
}

// CountEven returns the count of even numbers.
func CountEven(nums []int) int {
	// TODO
	return 0
}

// FindFirst returns index and value of the first item matching predicate.
// Return -1 and zero value when not found.
func FindFirst[T any](items []T, match func(T) bool) (int, T) {
	// TODO
	var zero T
	return -1, zero
}

// FizzBuzzRange prints numbers from 1 to n using FizzBuzz rules.
func FizzBuzzRange(n int) []string {
	// TODO: use a for loop
	return nil
}

// PrintPattern prints a right triangle of stars:
// *
// **
// ***
func PrintPattern(rows int) []string {
	// TODO
	return nil
}

// Example demonstrates labeled break.
func Example() {
	// TODO: create a nested loop and break both loops using a label
	fmt.Println("implement me")
}
