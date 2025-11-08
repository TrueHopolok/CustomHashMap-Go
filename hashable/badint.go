package hashable

type BadInt int

func (i BadInt) Hash(modulo uint64) uint64 {
	return uint64(i) % modulo
}

func (i BadInt) Equal(i2 BadInt) bool {
	return i == i2
}
