// hashmap desription TODO
package hashmap

//go:generate go run github.com/princjef/gomarkdoc/cmd/gomarkdoc . -o README.md

const (
	// MAX_CAPACITY that is allowed for allocation => max insertion capacity is 1 milliard.
	// That is around 1 gigabyte * sum of sizes of the key and value types.
	MAX_CAPACITY uint64 = 2_000_000_000

	// START_CAPACITY is how many elements are allowed to be inserted into default map created by [NewHashMap].
	START_CAPACITY uint64 = 64
)

type (
	// Hashable is interface implementing Hash function and comparison to be equal.
	// Hash function should be low-collision, otherwise time complexity may sky-rocket.
	//
	// For some type, Hashable interface is implemented in [hashable] package.
	Hashable interface {
		// Hash returns a numeric representation (hash) of the instance.
		// Returned hash must be in range of [0;modulo).
		Hash(modulo uint64) uint64

		// Equal returns true if 2 instances are equal, or false otherwise.
		Equal(other Hashable) bool
	}

	node[K Hashable, V any] struct {
		key  K           // key that relates to the value
		val  V           // data itself
		next *node[K, V] // stored as linked list, thus has field tpo the next element
	}

	// Item is used only by [All] function to return slice of all items.
	Item[K Hashable, V any] struct {
		Key K
		Val V
	}

	// HashMap is the main structure of the package. For indepth implementation, see this package's description.
	HashMap[K Hashable, V any] struct {
		capMax uint64 // maximum capacity of the hash table
		capCur uint64 // allowed capacity before expansion of the hash table
		len    uint64 // current amount of elements in the hash table

		table []*node[K, V] // hash table it self
	}
)
