package xxhash

/*
#include "C/xxhash.h"
*/
import "C"

import (
	"hash"
	"unsafe"
)

type xxHash64 struct {
	seed  uint64
	sum   uint64
	state unsafe.Pointer
}

// Size returns the number of bytes Sum will return.
func (xx *xxHash64) Size() int {
	return 8
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (xx *xxHash64) BlockSize() int {
	return 8
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (xx *xxHash64) Sum(in []byte) []byte {
	s := xx.Sum64()
	return append(in, byte(s>>56), byte(s>>48), byte(s>>40), byte(s>>32), byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func (xx *xxHash64) Write(p []byte) (n int, err error) {
	switch {
	case xx.state == nil:
		return 0, ErrAlreadyComputed
	case len(p) > oneGb:
		return 0, ErrMemoryLimit
	}
	C.XXH64_update(xx.state, unsafe.Pointer(&p[0]), C.uint(len(p)))
	return len(p), nil
}

func (xx *xxHash64) Sum64() uint64 {
	if xx.state == nil {
		return xx.sum
	}
	xx.sum = uint64(C.XXH64_digest(xx.state))
	xx.state = nil
	return xx.sum
}

// Reset resets the Hash to its initial state.
func (xx *xxHash64) Reset() {
	if xx.state != nil {
		C.XXH64_digest(xx.state)
	}
	xx.state = C.XXH64_init(C.ulonglong(xx.seed))
}

// NewS64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the specific seed.
func NewS64(seed uint64) hash.Hash64 {
	h := &xxHash64{
		seed: seed,
	}
	h.Reset()
	return h
}

// New64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the seed set to 0x0.
func New64() hash.Hash64 {
	return NewS64(0x0)
}

// Checksum64S returns the checksum of the input bytes with the specific seed.
func Checksum64S(in []byte, seed uint64) uint64 {
	return uint64(C.XXH64(unsafe.Pointer(&in[0]), C.uint(len(in)), C.ulonglong(seed)))
}

// Checksum64 returns the checksum of the input data with the seed set to 0
func Checksum64(in []byte) uint64 {
	return Checksum64S(in, 0x0)
}
