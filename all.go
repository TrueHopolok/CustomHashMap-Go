package hashmap

// All returns copy of all existing items in the [HashMap].
//
// Time Complexity:
//   - O(N), where N is current amount of elements
//
// Memory Complexity:
//   - O(N), where N is current amount of elements
func (m HashMap[K, V]) All() []Item[K, V] {
	if m.len == 0 {
		return nil
	}
	res := make([]Item[K, V], m.len)
	i := 0
	for _, n := range m.table {
		for n != nil {
			res[i] = Item[K, V]{
				Key: n.key,
				Val: n.val,
			}
			i++
		}
	}
	if i < len(res) {
		panic(IIError{"[ALL] collected different amount of items compared to stored amount in the map"})
	}
	return res
}

// Key returns copy of all existing keys in the [HashMap].
//
// Time Complexity:
//   - O(N), where N is current amount of elements
//
// Memory Complexity:
//   - O(N), where N is current amount of elements
func (m HashMap[K, V]) Key() []K {
	if m.len == 0 {
		return nil
	}
	res := make([]K, m.len)
	i := 0
	for _, n := range m.table {
		for n != nil {
			res[i] = n.key
			i++
		}
	}
	if i < len(res) {
		panic(IIError{"[KEY] collected different amount of items compared to stored amount in the map"})
	}
	return res
}

// Val returns copy of all existing values in the [HashMap].
//
// Time Complexity:
//   - O(N), where N is current amount of elements
//
// Memory Complexity:
//   - O(N), where N is current amount of elements
func (m HashMap[K, V]) Val() []V {
	if m.len == 0 {
		return nil
	}
	res := make([]V, m.len)
	i := 0
	for _, n := range m.table {
		for n != nil {
			res[i] = n.val
			i++
		}
	}
	if i < len(res) {
		panic(IIError{"[VAL] collected different amount of items compared to stored amount in the map"})
	}
	return res
}
