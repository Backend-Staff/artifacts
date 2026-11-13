package interfaces

import (
	"fmt"
	"strings"
)

// Shape is an interface with Area and Perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 {
	// TODO
	return 0
}

func (c Circle) Perimeter() float64 {
	// TODO
	return 0
}

type Rect struct{ Width, Height float64 }

func (r Rect) Area() float64 {
	// TODO
	return 0
}

func (r Rect) Perimeter() float64 {
	// TODO
	return 0
}

// TotalArea sums the area of all shapes.
func TotalArea(shapes []Shape) float64 {
	// TODO
	return 0
}

// Describe uses a type switch.
func Describe(value any) string {
	// TODO:
	// int -> "number"
	// string -> "text"
	// bool -> "boolean"
	// Shape -> "shape"
	// default -> "unknown"
	return ""
}

// Stringer demonstrates the standard fmt.Stringer interface.
type Person struct {
	First, Last string
}

func (p Person) String() string {
	// TODO
	return ""
}

// Greeter is a small interface.
type Greeter interface {
	Greet() string
}

// GreetAll calls Greet for every greeter.
func GreetAll(greeters []Greeter) []string {
	// TODO
	return nil
}

// English and Spanish implement Greeter.
type English struct{}
type Spanish struct{}

func (English) Greet() string {
	// TODO
	return ""
}

func (Spanish) Greet() string {
	// TODO
	return ""
}

// JoinUpper joins strings with ", " and uppercases the result.
func JoinUpper(parts []string) string {
	// TODO
	return strings.ToUpper(strings.Join(parts, ", "))
}

var _ = fmt.Sprintf
