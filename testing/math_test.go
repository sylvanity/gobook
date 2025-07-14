package math

import "testing"

// Basic test for the Add function
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

// Table-driven test for the Add function
func TestAddTableDriven(t *testing.T) {
	testCases := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -2, -3, -5},
		{"Mixed numbers", 2, -3, -1},
		{"Zero values", 0, 0, 0},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Mark subtest as safe to run in parallel
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// Benchmark for the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
