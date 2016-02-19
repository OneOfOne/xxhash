// +build !safe
// +build !appengine
// +build !be

package xxhash

import (
	"reflect"
	"unsafe"
)

// Backend returns the current version of xxhash being used.
const Backend = "GoUnsafe"

type byteReader struct {
	p unsafe.Pointer
}

func newbyteReader(b []byte) byteReader {
	return byteReader{unsafe.Pointer(&b[0])}
}

func (br byteReader) Uint32(i int) uint32 {
	return *(*uint32)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
}

func (br byteReader) Uint64(i int) uint64 {
	return *(*uint64)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
}

func (br byteReader) Byte(i int) byte {
	return *(*byte)(unsafe.Pointer(uintptr(br.p) + uintptr(i)))
}

func (xx *XXHash32) WriteString(s string) (int, error) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ss.Cap = ss.Len
	return xx.Write(*(*[]byte)(unsafe.Pointer(ss)))
}

func (xx *XXHash64) WriteString(s string) (int, error) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ss.Cap = ss.Len
	return xx.Write(*(*[]byte)(unsafe.Pointer(ss)))
}

func ChecksumString32(s string) uint32 {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ss.Cap = ss.Len
	return Checksum32(*(*[]byte)(unsafe.Pointer(ss)))
}

func ChecksumString32S(s string, seed uint32) uint32 {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ss.Cap = ss.Len
	return Checksum32S(*(*[]byte)(unsafe.Pointer(ss)), seed)
}

func ChecksumString64(s string) uint64 {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ss.Cap = ss.Len
	return Checksum64(*(*[]byte)(unsafe.Pointer(ss)))
}

func ChecksumString64S(s string, seed uint64) uint64 {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ss.Cap = ss.Len
	return Checksum64S(*(*[]byte)(unsafe.Pointer(ss)), seed)
}
