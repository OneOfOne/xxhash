// +build !appengine
// +build be

package xxhash

import "unsafe"

// Backend returns the current version of xxhash being used.
const Backend = "GoUnsafeBigEndian"

type byteReader struct {
	p unsafe.Pointer
}

func newbyteReader(b *[]byte) *byteReader {
	return &byteReader{unsafe.Pointer(&(*b)[0])}
}

func (br *byteReader) Uint32(i int) (u uint32) {
	u = *(*uint32)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
	if isBig {
		u = swap32be(u)
	}
	return
}

func (br *byteReader) Uint64(i int) (u uint64) {
	u = *(*uint64)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
	if isBig {
		u = swap64be(u)
	}
	return
}

func (br *byteReader) Byte(i int) byte {
	return *(*byte)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
}

func swap32be(x uint32) uint32 {
	return ((x << 24) & 0xff000000) |
		((x << 8) & 0x00ff0000) |
		((x >> 8) & 0x0000ff00) |
		((x >> 24) & 0x000000ff)
}

func swap64be(x uint64) uint64 {
	return ((x << 56) & 0xff00000000000000) |
		((x << 40) & 0x00ff000000000000) |
		((x << 24) & 0x0000ff0000000000) |
		((x << 8) & 0x000000ff00000000) |
		((x >> 8) & 0x00000000ff000000) |
		((x >> 24) & 0x0000000000ff0000) |
		((x >> 40) & 0x000000000000ff00) |
		((x >> 56) & 0x00000000000000ff)
}

var (
	dummy = [2]byte{1, 0}
	isBig = *(*int16)(unsafe.Pointer(&dummy[0])) != 1
)
