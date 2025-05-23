package main

import "testing"

func TestFibo(t *testing.T) {
	cases := []struct {
		n, want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2},
		{4, 3}, {5, 5}, {10, 55}, {15, 610},
	}

	for _, c := range cases {
		if got := fibo(c.n); got != c.want {
			t.Errorf("fibo(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}
