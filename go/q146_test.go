package program

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	tests := []struct {
		name string
	}{
		{
			name: "LRUCache",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.Put(1, 1)
			cache.Put(2, 2)
			fmt.Println(cache.Get(1))
			cache.Put(3, 3)
			fmt.Println(cache.Get(2))
			cache.Put(4, 4)
			cache.Put(4, 3)
			fmt.Println(cache.Get(1))
			fmt.Println(cache.Get(3))
			fmt.Println(cache.Get(4))
		})
	}
}
