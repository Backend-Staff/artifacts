package testing

import "testing"

func TestIsPrime(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{-1, false}, {0, false}, {1, false}, {2, true},
		{3, true}, {4, false}, {17, true}, {18, false},
	}

	for _, c := range cases {
		if got := IsPrime(c.in); got != c.want {
			t.Errorf("IsPrime(%d) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMedian(t *testing.T) {
	got, err := Median([]int{3, 1, 2})
	if err != nil {
		t.Fatal(err)
	}
	if got != 2 {
		t.Fatalf("Median() = %v, want 2", got)
	}

	if _, err := Median(nil); err == nil {
		t.Fatal("expected error for empty input")
	}
}
