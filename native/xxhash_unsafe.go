// +build !appengine
// +build !be

package xxhash

import "unsafe"

// Backend returns the current version of xxhash being used.
const Backend = "GoUnsafe"

type byteReader struct {
	p unsafe.Pointer
}

func newbyteReader(b []byte) *byteReader {
	return &byteReader{unsafe.Pointer(&b[0])}
}

func (br *byteReader) Uint32(i int) (u uint32) {
	u = *(*uint32)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
	return
}

func (br *byteReader) Uint64(i int) (u uint64) {
	u = *(*uint64)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
	return
}

func (br *byteReader) Byte(i int) byte {
	return *(*byte)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
}
