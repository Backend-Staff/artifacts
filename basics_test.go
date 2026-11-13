package basics

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(1, 2, 3); got != 6 {
		t.Fatalf("Sum() = %d, want 6", got)
	}
}

func TestReverse(t *testing.T) {
	cases := []struct{ in, want string }{
		{"abc", "cba"},
		{"héllo", "olléh"},
	}
	for _, c := range cases {
		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := map[int]string{1: "1", 3: "Fizz", 5: "Buzz", 15: "FizzBuzz"}
	for in, want := range cases {
		if got := FizzBuzz(in); got != want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 1, 3, 2})
	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
