package hashable

import (
	hashmap "github.com/TrueHopolok/CustomHashMap-Go"
)

type Int int

func (n Int) Hash(modulo uint64) uint64 {
	d := make([]byte, 8)
	for i := 0; i < 8; i++ {
		d[i] = byte(n >> (8 * i))
	}
	return FNV1(d) % modulo
}

func (n Int) Equal(n2 hashmap.Hashable) bool {
	v2, ok := n2.(Int)
	if !ok {
		panic("invalid type comparison in hashable.Int")
	}
	return n == v2
}
