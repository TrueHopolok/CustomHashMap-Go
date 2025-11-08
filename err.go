package hashmap

import "fmt"

// All errors are thrown by panic and are not returned as values.
type (
	// HashableError indicate that provided hash function is returning value out of given 0 to modulo range.
	HashableError struct {
		keyHashed uint64
		modulo    uint64
	}

	// IIError - Internal Implementation Error, something happend that should have been based on implementation.
	IIError struct {
		msg string
	}

	// MaxCapExceededError indicate that capacity that was / will be requested is too large being larger than [MAX_CAPACITY].
	MaxCapExceededError struct{}

	// ZeroLengthError occurs only in [MakeHashMap] if given length is zero or larger than [MAX_CAPACITY] divided by 2.
	InvalidLengthError struct{}
)

func (e HashableError) Error() string {
	return fmt.Sprintf("given key returned hash [%d] which is out of range [0:%d)", e.keyHashed, e.modulo)
}

func (e IIError) Error() string {
	return e.msg
}

func (e MaxCapExceededError) Error() string {
	return "new request capacity exceeds MAX_CAPACITY value"
}

func (e InvalidLengthError) Error() string {
	return "given length is unacceptable being 0 or too large"
}
