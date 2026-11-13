package structs

import (
	"errors"
	"fmt"
	"time"
)

// User represents a basic user.
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// NewUser creates a user with CreatedAt set to time.Now().
func NewUser(id int, name, email string) (*User, error) {
	// TODO: validate name and email are non-empty
	return nil, errors.New("not implemented")
}

// String implements fmt.Stringer.
func (u User) String() string {
	// TODO: return "Ada <ada@example.com>"
	return ""
}

// Rename changes the user's name.
func (u *User) Rename(name string) error {
	// TODO
	return nil
}

// Address is embedded into Customer.
type Address struct {
	City    string
	Country string
}

type Customer struct {
	User
	Address
	LoyaltyPoints int
}

// FullDescription returns "Ada lives in London, United Kingdom".
func (c Customer) FullDescription() string {
	// TODO
	return ""
}

// Rectangle demonstrates methods with value and pointer receivers.
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	// TODO
	return 0
}

func (r *Rectangle) Scale(factor float64) error {
	// TODO: reject factor <= 0
	return nil
}

// Stack is a generic LIFO stack.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	// TODO
}

func (s *Stack[T]) Pop() (T, error) {
	// TODO: return error when empty
	var zero T
	return zero, nil
}

func (s *Stack[T]) Len() int {
	// TODO
	return 0
}

// ExampleFormatting shows struct formatting.
func ExampleFormatting() {
	u := User{ID: 1, Name: "Ada"}
	fmt.Printf("%+v\n", u)
}
