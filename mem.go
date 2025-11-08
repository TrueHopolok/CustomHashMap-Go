package hashmap

// expand multiple capacity of the map by 2
func (m HashMap[K, V]) expand() {
	items := m.All()

	for {
		m.capCur *= 2

		if m.capCur*2 > MAX_CAPACITY {
			panic(MaxCapExceededError{})
		}

		if m.len <= m.capCur {
			break
		}
	}

	m.capMax = m.capCur * 2

	m.len = 0
	m.table = make([]*node[K, V], m.capMax)
	for _, item := range items {
		m.Add(item.Key, item.Val)
	}

	if m.len != uint64(len(items)) {
		panic(IIError{"[EXPAND] inserted different amount of elements that previously it was"})
	}
}
