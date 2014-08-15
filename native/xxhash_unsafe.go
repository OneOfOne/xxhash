// +build !appengine

package xxhash

import "unsafe"

// Backend returns the current version of xxhash being used.
const Backend = "GoUnsafe"

func readU32le(b []byte, i int) uint32 {
	return *(*uint32)(unsafe.Pointer(&b[i]))
}

func readU64le(b []byte, i int) uint64 {
	return *(*uint64)(unsafe.Pointer(&b[i]))
}
