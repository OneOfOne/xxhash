// +build !safe
// +build !appengine

package xxhash

import (
	"io"
	"reflect"
	"unsafe"
)

// Backend returns the current version of xxhash being used.
const Backend = "GoUnsafe"

// ChecksumString32S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString32S(s string, seed uint32) uint32 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum32S((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

// ChecksumString64S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString64S(s string, seed uint64) uint64 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum64S((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

func writeString(w io.Writer, s string) (int, error) {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return w.Write((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)])
}
