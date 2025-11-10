package hashable

const (
	FNV1_basis uint64 = 14695981039346656037
	FNV1_prime uint64 = 1099511628211
)

// FNV1 returns hash of a bytes slice as 64bit number.
// Any data that is encoded may be passed.
//
// Algorithm and its description:
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1_hash
func FNV1(data []byte) uint64 {
	hash := FNV1_basis
	for _, b := range data {
		hash *= FNV1_prime
		hash ^= uint64(b)
	}
	return hash
}
