package hashmap_test

import (
	"fmt"

	hashmap "github.com/TrueHopolok/CustomHashMap-Go"
	"github.com/TrueHopolok/CustomHashMap-Go/hashable"
)

func ExampleHashMap() {
	input := "ABCABA"
	hmap := hashmap.NewHashMap[hashable.Rune, int]()
	for _, r := range input {
		hr := hashable.Rune(r) // typecast required
		v := hmap.Req(hr) + 1
		hmap.Set(hr, v)
	}
	results := hmap.All()
	for _, item := range results {
		fmt.Printf("%c: %d\n", item.Key, item.Val)
	}
	// Output:
	// A: 3
	// C: 1
	// B: 2
}
