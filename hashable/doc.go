// hashable is sub-pacakge for [github.com/TrueHopolok/CustomHashMap-Go]
// that implements [github.com/TrueHopolok/CustomHashMap-Go.Hashable] interface for basic types.
//
// Package uses FNV-1 hash function to encoded any data.
// This function can be used to hash any other type.
// For more info about this type, see [FNV1].
package hashable

//go:generate go run github.com/princjef/gomarkdoc/cmd/gomarkdoc . -o README.md
