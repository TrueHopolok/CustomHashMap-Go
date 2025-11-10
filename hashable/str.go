package hashable

import hashmap "github.com/TrueHopolok/CustomHashMap-Go"

type Str string

func (s Str) Hash(modulo uint64) uint64 {
	return FNV1([]byte(s)) % modulo
}

func (s Str) Equal(s2 hashmap.Hashable) bool {
	v2, ok := s2.(Str)
	if !ok {
		panic("invalid type comparison in hashable.Str")
	}
	if s != s2 {
		return false
	}
	for i := range v2 {
		if s[i] != v2[i] {
			return false
		}
	}
	return true
}
