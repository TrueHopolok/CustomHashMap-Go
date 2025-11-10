package hashable

import (
	hashmap "github.com/TrueHopolok/CustomHashMap-Go"
)

type Rune rune

func (n Rune) Hash(modulo uint64) uint64 {
	d := make([]byte, 4)
	for i := 0; i < 4; i++ {
		d[i] = byte(n >> (8 * i))
	}
	return FNV1(d) % modulo
}

func (n Rune) Equal(n2 hashmap.Hashable) bool {
	v2, ok := n2.(Rune)
	if !ok {
		panic("invalid type comparison in hashable.Int")
	}
	return n == v2
}
