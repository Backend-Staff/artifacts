package generics

import "constraints"

// Number is a constraint for numeric types.
type Number interface {
	constraints.Integer | constraints.Float
}

// Max returns the larger value.
func Max[T constraints.Ordered](a, b T) T {
	// TODO
	return a
}

// Min returns the smaller value.
func Min[T constraints.Ordered](a, b T) T {
	// TODO
	return a
}

// Map transforms each item.
func Map[T, R any](items []T, fn func(T) R) []R {
	// TODO
	return nil
}

// Keys returns map keys.
func Keys[K comparable, V any](m map[K]V) []K {
	// TODO
	return nil
}

// Values returns map values.
func Values[K comparable, V any](m map[K]V) []V {
	// TODO
	return nil
}

// Sum adds numeric values.
func Sum[T Number](items []T) T {
	// TODO
	return 0
}

// Pair is a generic key-value pair.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// ToPairs converts a map to pairs.
func ToPairs[K comparable, V any](m map[K]V) []Pair[K, V] {
	// TODO
	return nil
}
