package cache

import (
	"runtime"
	"sync"
	"testing"
)

func TestCache_GetSet(t *testing.T) {
	c := New[string, int]()
	c.Set("answer", 42)

	if v, ok := c.Get("answer"); !ok || v != 42 {
		t.Fatalf("want 42, got %v (ok=%v)", v, ok)
	}

	if _, ok := c.Get("missing"); ok {
		t.Fatalf("expected missing key to return ok=false")
	}
}

func TestCache_Concurrent(t *testing.T) {
	c := New[int, int]()
	const n = 1_000

	var wg sync.WaitGroup
	wg.Add(2 * n)

	// writers
	for i := 0; i < n; i++ {
		go func(i int) {
			c.Set(i, i*i)
			wg.Done()
		}(i)
	}

	// readers
	for i := 0; i < n; i++ {
		go func(i int) {
			_, _ = c.Get(i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	runtime.GC()
}
