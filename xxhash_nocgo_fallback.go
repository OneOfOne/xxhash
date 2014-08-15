// +build !cgo

package xxhash

import (
	"hash"

	N "github.com/OneOfOne/xxhash/native"
)

// Backend returns the current version of xxhash being used.
const Backend = N.Backend

// NewS32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting with the specific seed.
func NewS32(seed uint32) hash.Hash32 {
	return N.NewS32(0)
}

// New32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting with the seed set to 0x0.
func New32() hash.Hash32 {
	return NewS32(0x0)
}

// NewS64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the specific seed.
func NewS64(seed uint64) hash.Hash64 {
	return N.NewS64(seed)
}

// New64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the seed set to 0x0.
func New64() hash.Hash64 {
	return NewS64(0x0)
}

// Checksum64S returns the checksum of the input bytes with the specific seed.
func Checksum64S(in []byte, seed uint64) uint64 {
	return N.Checksum64S(in, seed)
}

// Checksum64 returns the checksum of the input data with the seed set to 0
func Checksum64(in []byte) uint64 {
	return Checksum64S(in, 0x0)
}

// Checksum32S returns the checksum of the input bytes with the specific seed.
func Checksum32S(in []byte, seed uint32) uint32 {
	return N.Checksum32S(in, seed)
}

// Checksum32 returns the checksum of the input data with the seed set to 0
func Checksum32(in []byte) uint32 {
	return Checksum32S(in, 0x0)
}
