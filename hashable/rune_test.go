package hashable_test

import (
	"fmt"

	hashmap "github.com/TrueHopolok/CustomHashMap-Go"
	"github.com/TrueHopolok/CustomHashMap-Go/hashable"
)

func ExampleRune() {
	m := hashmap.NewHashMap[hashable.Rune, string]()
	m.Set('h', "Hello")
	m.Set('w', "World")
	fmt.Printf("%s_%s!", m.Req('h'), m.Req('w'))
	// Output:
	// Hello_World!
}
