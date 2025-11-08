package hashmap

// NewHashMap creates new empty [HashMap] with reserved memory of base capacity.
func NewHashMap[K Hashable, V any]() HashMap[K, V] {
	return HashMap[K, V]{
		capMax: START_CAPACITY * 2,
		capCur: START_CAPACITY,
		len:    0,
		table:  make([]*node[K, V], START_CAPACITY*2),
	}
}

// MakeHashMap creates new [HashMap] with reserved memory for given len parameter.
// Gurantees no resizing after calling this function until amount of inserted elements surpasses given length.
// Beaware, total memory allocated will be 2 times larger than given len to avoid collisions.
//
// Time Complexity:
//   - O(N), where N is given length
//
// Memory Complexity:
//   - O(2N), where N is given length
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
