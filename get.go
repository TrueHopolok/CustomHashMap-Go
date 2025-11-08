package hashmap

// Get returns value by given key and boolean indicating if such item exists.
// If item with given key does not exists, value will be initialized as default value and false will be returned.
//
// Time Complexity:
//   - O(1), if [Hashable.Hash] is low-collision
//
// Memory Complexity:
//   - O(1)
func (m HashMap[K, V]) Get(key K) (val V, exists bool) {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	cur := m.table[keyHashed]
	if cur == nil {
		return val, false
	}
	for cur != nil {
		if key.Equal(cur.key) {
			return cur.val, true
		}
		cur = cur.next
	}
	return val, false
}

// Has returns whether or not item with given key exists.
//
// Time Complexity:
//   - O(1), if [Hashable.Hash] is low-collision
//
// Memory Complexity:
//   - O(1)
func (m HashMap[K, V]) Has(key K) (exists bool) {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	cur := m.table[keyHashed]
	if cur == nil {
		return false
	}
	for cur != nil {
		if key.Equal(cur.key) {
			return true
		}
		cur = cur.next
	}
	return false
}

// Req returns value of the item with given key or default value of the type if item with given key does not exists.
//
// Time Complexity:
//   - O(1), if [Hashable.Hash] is low-collision
//
// Memory Complexity:
//   - O(1)
func (m HashMap[K, V]) Req(key K) (val V) {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	cur := m.table[keyHashed]
	if cur == nil {
		return val
	}
	for cur != nil {
		if key.Equal(cur.key) {
			return cur.val
		}
		cur = cur.next
	}
	return val
}
