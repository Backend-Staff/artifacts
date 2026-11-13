package projects

import "testing"

func TestInventory(t *testing.T) {
	inv := NewInventory()

	id, err := inv.Add("Keyboard", 49.99, 10)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.UpdateStock(id, -3); err != nil {
		t.Fatal(err)
	}

	p, err := inv.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if p.Stock != 7 {
		t.Fatalf("stock = %d, want 7", p.Stock)
	}

	if err := inv.UpdateStock(id, -100); err == nil {
		t.Fatal("expected error for negative stock")
	}
}
