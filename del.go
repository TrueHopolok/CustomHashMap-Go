package hashmap

// Del removes value with given key from the map.
// If such element does not exists, does nothing.
//   - To gurantee deletion of existing element, use [Rem].
//
// Time Complexity:
//   - O(1), if [Hashable.Hash] is low-collision
//
// Memory Complexity:
//   - O(1)
func (m *HashMap[K, V]) Del(key K) {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	cur := m.table[keyHashed]
	if cur == nil {
		return
	}
	if key.Equal(cur.key) {
		m.table[keyHashed] = nil
		m.len--
		return
	}
	prev := cur
	cur = cur.next
	for cur != nil {
		if key.Equal(cur.key) {
			prev.next = cur.next
			m.len--
			return
		}
		cur, prev = cur.next, cur
	}
}

// Rem removes value with given key from the map.
// Returns whether or not item was deleted.
//   - To gurantee deletion of the key without checking its existance, use [Del].
//
// Time Complexity:
//   - O(1), if [Hashable.Hash] is low-collision
//
// Memory Complexity:
//   - O(1)
func (m *HashMap[K, V]) Rem(key K) bool {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	cur := m.table[keyHashed]
	if cur == nil {
		return false
	}
	if key.Equal(cur.key) {
		m.table[keyHashed] = nil
		m.len--
		return true
	}
	prev := cur
	cur = cur.next
	for cur != nil {
		if key.Equal(cur.key) {
			prev.next = cur.next
			m.len--
			return true
		}
		cur, prev = cur.next, cur
	}
	return false
}
