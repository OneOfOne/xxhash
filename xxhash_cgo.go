// +build cgo
// +build forcecgo !go1.7

package xxhash

//go:generate git subtree pull --prefix vendor/xxHash https://github.com/Cyan4973/xxHash master --squash

/*
#cgo CFLAGS: -O3 -std=c99 -pedantic
#include "vendor/xxHash/xxhash.c"
*/
import "C"

import (
	"hash"
	"reflect"
	"unsafe"
)

// Backend returns the current version of xxhash being used.
const Backend = "CGO"

type XXHash32 struct {
	seed  uint32
	state C.XXH32_state_t
}

var _ interface {
	hash.Hash32
	WriteString(string) (int, error)
} = (*XXHash32)(nil)

// Size returns the number of bytes Sum will return.
func (xx *XXHash32) Size() int {
	return 4
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (xx *XXHash32) BlockSize() int {
	return 16
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (xx *XXHash32) Sum(in []byte) []byte {
	s := xx.Sum32()
	return append(in, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func (xx *XXHash32) Write(p []byte) (n int, err error) {
	C.XXH32_update(&xx.state, unsafe.Pointer(&p[0]), C.size_t(len(p)))
	return len(p), nil
}

func (xx *XXHash32) WriteString(s string) (int, error) {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return xx.Write((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)])
}

func (xx *XXHash32) Sum32() uint32 {
	return uint32(C.XXH32_digest(&xx.state))
}

// Reset resets the Hash to its initial state.
func (xx *XXHash32) Reset() {
	C.XXH32_reset(&xx.state, C.uint(xx.seed))
}

// NewS32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting with the specific seed.
func NewS32(seed uint32) *XXHash32 {
	h := &XXHash32{
		seed: seed,
	}
	h.Reset()
	return h
}

// New32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting with the seed set to 0.
func New32() *XXHash32 {
	return NewS32(0)
}

// Checksum32S returns the checksum of the input data with the specific seed.
func Checksum32S(in []byte, seed uint32) uint32 {
	return uint32(C.XXH32(unsafe.Pointer(&in[0]), C.size_t(len(in)), C.uint(seed)))
}

// Checksum32 returns the checksum of the input data with the seed set to 0
func Checksum32(in []byte) uint32 {
	return Checksum32S(in, 0)
}

// ChecksumString32 returns the checksum of the input data, without creating a copy, with the seed set to 0.
func ChecksumString32(s string) uint32 {
	return ChecksumString32S(s, 0)
}

// ChecksumString32S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString32S(s string, seed uint32) uint32 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum32S((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

type XXHash64 struct {
	seed  uint64
	state C.XXH64_state_t
}

var _ interface {
	hash.Hash64
	WriteString(string) (int, error)
} = (*XXHash64)(nil)

// Size returns the number of bytes Sum will return.
func (xx *XXHash64) Size() int {
	return 8
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (xx *XXHash64) BlockSize() int {
	return 32
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (xx *XXHash64) Sum(in []byte) []byte {
	s := xx.Sum64()
	return append(in, byte(s>>56), byte(s>>48), byte(s>>40), byte(s>>32), byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func (xx *XXHash64) Write(p []byte) (n int, err error) {
	C.XXH64_update(&xx.state, unsafe.Pointer(&p[0]), C.size_t(len(p)))
	return len(p), nil
}

func (xx *XXHash64) WriteString(s string) (int, error) {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return xx.Write((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)])
}

func (xx *XXHash64) Sum64() uint64 {
	return uint64(C.XXH64_digest(&xx.state))
}

// Reset resets the Hash to its initial state.
func (xx *XXHash64) Reset() {
	C.XXH64_reset(&xx.state, C.ulonglong(xx.seed))

}

// NewS64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the specific seed.
func NewS64(seed uint64) *XXHash64 {
	h := &XXHash64{
		seed: seed,
	}
	h.Reset()
	return h
}

// New64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the seed set to 0.
func New64() *XXHash64 {
	return NewS64(0)
}

// Checksum64S returns the checksum of the input bytes with the specific seed.
func Checksum64S(in []byte, seed uint64) uint64 {
	return uint64(C.XXH64(unsafe.Pointer(&in[0]), C.size_t(len(in)), C.ulonglong(seed)))
}

// Checksum64 returns the checksum of the input data with the seed set to 0.
func Checksum64(in []byte) uint64 {
	return Checksum64S(in, 0)
}

// ChecksumString64 returns the checksum of the input data, without creating a copy, with the seed set to 0.
func ChecksumString64(s string) uint64 {
	return ChecksumString64S(s, 0)
}

// ChecksumString64S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString64S(s string, seed uint64) uint64 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum64S((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}
