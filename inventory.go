package projects

import (
	"errors"
	"sort"
	"sync"
)

// Product is an inventory item.
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Inventory is a thread-safe in-memory inventory.
type Inventory struct {
	mu       sync.RWMutex
	products map[int]Product
	nextID   int
}

// NewInventory creates an empty inventory.
func NewInventory() *Inventory {
	return &Inventory{products: make(map[int]Product)}
}

// Add creates a product and returns its ID.
func (i *Inventory) Add(name string, price float64, stock int) (int, error) {
	// TODO: validate name, price >= 0, stock >= 0
	return 0, nil
}

// Get returns a copy of a product.
func (i *Inventory) Get(id int) (Product, error) {
	// TODO: return ErrNotFound when missing
	return Product{}, nil
}

// UpdateStock changes stock by delta, preventing negative stock.
func (i *Inventory) UpdateStock(id, delta int) error {
	// TODO
	return nil
}

// List returns products sorted by name.
func (i *Inventory) List() []Product {
	// TODO
	return nil
}

// TotalValue returns sum(price * stock).
func (i *Inventory) TotalValue() float64 {
	// TODO
	return 0
}

var (
	ErrNotFound      = errors.New("product not found")
	ErrInvalidInput  = errors.New("invalid input")
)

var _ = sort.Slice
