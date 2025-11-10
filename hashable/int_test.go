package hashable_test

import (
	"fmt"

	hashmap "github.com/TrueHopolok/CustomHashMap-Go"
	"github.com/TrueHopolok/CustomHashMap-Go/hashable"
)

func ExampleInt() {
	m := hashmap.NewHashMap[hashable.Int, string]()
	m.Set(2, "Hello")
	m.Set(5, "World")
	fmt.Printf("%s_%s!", m.Req(2), m.Req(5))
	// Output:
	// Hello_World!
}
