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
	return Checksum32S((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s)], seed)
}

// ChecksumString64S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString64S(s string, seed uint64) uint64 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum64S((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s)], seed)
}

func writeString(w io.Writer, s string) (int, error) {
	if len(s) == 0 {
		return w.Write(nil)
	}
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return w.Write((*[0x7fffffff]byte)(unsafe.Pointer(ss.Data))[:len(s)])
}

// Checksum64S returns the 64bit xxhash checksum for a single input
func checksum64S(in []byte, seed uint64) uint64 {
	var (
		wordsLen = int(uint32(len(in)) >> 3)
		words    = (*(*[]uint64)(unsafe.Pointer(&in)))[:wordsLen:wordsLen]

		h  uint64 = prime64x5
		v1        = seed + prime64x1 + prime64x2
		v2        = seed + prime64x2
		v3        = seed + 0
		v4        = seed - prime64x1

		i int
	)

	for ; i < len(words)-3; i += 4 {
		v1 += words[i] * prime64x2
		v1 = rotl64_31(v1) * prime64x1

		v2 += words[i+1] * prime64x2
		v2 = rotl64_31(v2) * prime64x1

		v3 += words[i+2] * prime64x2
		v3 = rotl64_31(v3) * prime64x1

		v4 += words[i+3] * prime64x2
		v4 = rotl64_31(v4) * prime64x1
	}

	h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)

	h = doVx64(h, v1)
	h = doVx64(h, v2)
	h = doVx64(h, v3)
	h = doVx64(h, v4)

	h += uint64(len(in))

	for _, k := range words[i:] {
		k *= prime64x2
		k = rotl64_31(k)
		k *= prime64x1
		h ^= k
		h = rotl64_27(h)*prime64x1 + prime64x4
		i++
	}

	if i = (i << 3); i < len(in)-3 {
		in := in[i : i+4 : len(in)]
		h ^= uint64(u32(in[0:4:len(in)])) * prime64x1
		h = rotl64_23(h)*prime64x2 + prime64x3

		i += 4
	}

	for ; i < len(in); i++ {
		h ^= uint64(in[i]) * prime64x5
		h = rotl64_11(h) * prime64x1
	}

	return mix64(h)
}
