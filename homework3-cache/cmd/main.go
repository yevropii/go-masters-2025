package main

import (
	"fmt"
	"github.com/yevropii/go-masters-2025/homework3-cache/internal/cache"
)

func main() {
	c := cache.New[string, int]()
	c.Set("one", 1)
	c.Set("two", 2)

	if v, ok := c.Get("one"); ok {
		fmt.Println("two = ", v)
	}
}
