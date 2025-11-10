package hashmap

// Set replaces value of item in the [HashMap] with given key.
// May increase size of the [HashMap].
//
// Items with hash of the key are stored in linked list,
// so any happend collision is handled by inserting item into the end of the linked list.
//
//   - To gurantee insertion of the value as a new item, use [Add].
//   - To gurantee replacement of the old value, use [Upd].
//
// Time Complexity:
//   - O(1) amortized, if [Hashable.Hash] is low-collision
//   - O(N) worst, where N is current amount of elements
//
// Memory Complexity:
//   - O(1) amortized
//   - O(N) worst, where N is current amount of elements
func (m *HashMap[K, V]) Set(key K, val V) {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	if m.Has(key) {
		cur := m.table[keyHashed]
		for cur != nil {
			if key.Equal(cur.key) {
				cur.val = val
				return
			}
			cur = cur.next
		}
		panic(IIError{"[SET] haven't found the key while [HAS] found it previously"})
	}

	for m.len >= m.capCur {
		m.expand()
		keyHashed = key.Hash(m.capMax)
		if keyHashed >= m.capMax {
			panic(HashableError{keyHashed, m.capMax})
		}
	}
	m.len++

	cur := m.table[keyHashed]
	if cur == nil {
		m.table[keyHashed] = &node[K, V]{
			key:  key,
			val:  val,
			next: nil,
		}
		return
	}
	for cur != nil {
		if key.Equal(cur.key) {
			panic(IIError{"[SET] key was found even though, [HAS] haven't found it previously"})
		}
		if cur.next == nil {
			cur.next = &node[K, V]{
				key:  key,
				val:  val,
				next: nil,
			}
			return
		}
		cur = cur.next
	}
	panic(IIError{"[SET] execution ended without updating any item"})
}

// Add inserts value into the [HashMap] as item with given key.
// Returns whether or not value was inserted into the [HashMap].
// May increase size of the [HashMap].
//
// Items with hash of the key are stored in linked list,
// so any happend collision is handled by inserting item into the end of the linked list.
//
//   - To gurantee value getting saved into the item (new or existing) with given key, use [Set].
//   - To gurantee replacement of the old value, use [Upd].
//
// Time Complexity:
//   - O(1) amortized, if [Hashable.Hash] is low-collision
//   - O(N) worst, where N is current amount of elements
//
// Memory Complexity:
//   - O(1) amortized
//   - O(N) worst, where N is current amount of elements
func (m *HashMap[K, V]) Add(key K, val V) bool {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	if m.Has(key) {
		return false
	}

	for m.len >= m.capCur {
		m.expand()
		keyHashed = key.Hash(m.capMax)
		if keyHashed >= m.capMax {
			panic(HashableError{keyHashed, m.capMax})
		}
	}
	m.len++

	cur := m.table[keyHashed]
	if cur == nil {
		m.table[keyHashed] = &node[K, V]{
			key:  key,
			val:  val,
			next: nil,
		}
		return true
	}
	for cur != nil {
		if key.Equal(cur.key) {
			panic(IIError{"[ADD] key was found even though, [HAS] haven't found it previously"})
		}
		if cur.next == nil {
			cur.next = &node[K, V]{
				key:  key,
				val:  val,
				next: nil,
			}
			return true
		}
		cur = cur.next
	}
	panic(IIError{"[ADD] execution ended without updating any item"})
}

// Upd updates value of item in the [HashMap] with given key.
// Returns whether or not value of the item was found and updated.
// Never increases size of the [HashMap].
//
//   - To gurantee value getting saved into the item (new or existing) with given key, use [Set].
//   - To gurantee insertion of the value as a new item, use [Add].
//
// Time Complexity:
//   - O(1) amortized, if [Hashable.Hash] is low-collision
//   - O(N) worst, where N is current amount of elements
//
// Memory Complexity:
//   - O(1) amortized
//   - O(N) worst, where N is current amount of elements
func (m *HashMap[K, V]) Upd(key K, val V) bool {
	keyHashed := key.Hash(m.capMax)
	if keyHashed >= m.capMax {
		panic(HashableError{keyHashed, m.capMax})
	}

	if m.Has(key) {
		cur := m.table[keyHashed]
		for cur != nil {
			if key.Equal(cur.key) {
				cur.val = val
				return true
			}
			cur = cur.next
		}
		panic(IIError{"[UPD] haven't found the key while [HAS] found it previously"})
	}

	return false
}
