// +build !safe
// +build !appengine

package xxhash

import (
	"reflect"
	"unsafe"
)

// Backend returns the current version of xxhash being used.
const Backend = "GoUnsafe"

// ChecksumString32S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString32S(s string, seed uint32) uint32 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum32S((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

// ChecksumString64S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString64S(s string, seed uint64) uint64 {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum64S((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

func writeString32(w *XXHash32, s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return w.Write((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)])
}

func writeString64(w *XXHash64, s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return w.Write((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s)])
}

// Checksum64S returns the 64bit xxhash checksum for a single input
func checksum64S(in []byte, seed uint64) uint64 {
	var (
		wordsLen = int(uint32(len(in)) >> 3)
		words    = ((*[maxInt32]uint64)(unsafe.Pointer(&in[0])))[:wordsLen:wordsLen]

		h uint64 = prime64x5

		v1 = seed + prime64x1 + prime64x2
		v2 = seed + prime64x2
		v3 = seed + 0
		v4 = seed - prime64x1

		i int
	)

	for m := len(words) - 3; i < m; i += 4 {
		_ = words[i+3]

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

	for ; i < len(words); i++ {
		k := words[i]
		k *= prime64x2
		k = rotl64_31(k)
		k *= prime64x1
		h ^= k
		h = rotl64_27(h)*prime64x1 + prime64x4
	}

	if i = (i << 3); i < len(in)-3 {
		h ^= uint64(((*[1]uint32)(unsafe.Pointer(&in[i])))[0]) * prime64x1
		h = rotl64_23(h)*prime64x2 + prime64x3

		i += 4
	}

	for _, b := range in[i:] {
		h ^= uint64(b) * prime64x5
		h = rotl64_11(h) * prime64x1
	}

	return mix64(h)
}
