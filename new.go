package hashmap

// NewHashMap creates new empty [HashMap] with reserved memory of base capacity.
func NewHashMap[K Hashable, V any]() HashMap[K, V] {
	return HashMap[K, V]{
		capMax: 2,
		capCur: 1,
		len:    0,
		table:  make([]*node[K, V], 2),
	}
}

// MakeHashMap creates new [HashMap] with reserved memory for given len parameter.
// Gurantees no resizing after calling this function until amount of inserted elements surpasses given length.
// Beaware, total memory allocated will be 2 times larger than given len to avoid collisions.
//
// Parameter must be between 1 and [MAX_CAPACITY]/2 included both points, otherwise code will panic.
//
// Time Complexity:
//   - O(N), where N is given length
//
// Memory Complexity:
//   - O(N), where N is given length
func MakeHashMap[K Hashable, V any](len uint64) HashMap[K, V] {
	if len == 0 || len > MAX_CAPACITY/2 {
		panic(InvalidLengthError{})
	}
	return HashMap[K, V]{
		capMax: len * 2,
		capCur: len,
		len:    0,
		table:  make([]*node[K, V], len*2),
	}
}
