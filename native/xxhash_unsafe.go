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

type byteReader struct {
	p unsafe.Pointer
}

func newbyteReader(b []byte) byteReader {
	if len(b) == 0 {
		return byteReader{}
	}
	return byteReader{unsafe.Pointer(&b[0])}
}

func (br byteReader) Uint32(i int32) uint32 {
	u := *(*uint32)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
	if isBigEndian {
		u = swap32(u)
	}
	return u
}

func (br byteReader) Uint64(i int32) uint64 {
	u := *(*uint64)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
	if isBigEndian {
		u = swap64(u)
	}
	return u
}

func (br byteReader) Byte(i int32) byte {
	return *(*byte)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
}

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

func swap32(x uint32) uint32 {
	return ((x << 24) & 0xff000000) |
		((x << 8) & 0x00ff0000) |
		((x >> 8) & 0x0000ff00) |
		((x >> 24) & 0x000000ff)
}

func swap64(x uint64) uint64 {
	x = (0xff00ff00ff00ff & (x >> 8)) | ((0xff00ff00ff00ff & x) << 8)
	x = (0xffff0000ffff & (x >> 16)) | ((0xffff0000ffff & x) << 16)
	return (0xffffffff & (x >> 32)) | ((0xffffffff & x) << 32)
}
