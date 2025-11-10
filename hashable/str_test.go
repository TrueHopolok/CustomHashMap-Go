package hashable_test

import (
	"fmt"

	hashmap "github.com/TrueHopolok/CustomHashMap-Go"
	"github.com/TrueHopolok/CustomHashMap-Go/hashable"
)

func ExampleStr() {
	m := hashmap.NewHashMap[hashable.Str, string]()
	m.Set("hi", "Hello")
	m.Set("www", "World")
	fmt.Printf("%s_%s!", m.Req("hi"), m.Req("www"))
	// Output:
	// Hello_World!
}
