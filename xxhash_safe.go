// +build appengine safe

package xxhash

import "io"

// Backend returns the current version of xxhash being used.
const Backend = "GoSafe"

func ChecksumString32S(s string, seed uint32) uint32 {
	return Checksum32S([]byte(s), seed)
}

func ChecksumString64S(s string, seed uint64) uint64 {
	return Checksum64S([]byte(s), seed)
}

func writeString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(s))
}

func checksum64S(in []byte, seed uint64) (h uint64) {
	var (
		v1 = seed + prime64x1 + prime64x2
		v2 = seed + prime64x2
		v3 = seed + 0
		v4 = seed - prime64x1

		i int
	)

	for ; i < len(in)-31; i += 32 {
		in := in[i : i+32 : len(in)]
		v1 += u64(in[0:8:len(in)]) * prime64x2
		v1 = rotl64_31(v1) * prime64x1

		v2 += u64(in[8:16:len(in)]) * prime64x2
		v2 = rotl64_31(v2) * prime64x1

		v3 += u64(in[16:24:len(in)]) * prime64x2
		v3 = rotl64_31(v3) * prime64x1

		v4 += u64(in[24:32:len(in)]) * prime64x2
		v4 = rotl64_31(v4) * prime64x1
	}

	h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)

	h = doVx64(h, v1)
	h = doVx64(h, v2)
	h = doVx64(h, v3)
	h = doVx64(h, v4)

	h += uint64(len(in))

	for ; i < len(in)-7; i += 8 {
		k := u64(in[i : i+8 : len(in)])
		k *= prime64x2
		k = rotl64_31(k)
		k *= prime64x1
		h ^= k
		h = rotl64_27(h)*prime64x1 + prime64x4
	}

	for ; i < len(in)-3; i += 4 {
		h ^= uint64(u32(in[i:i+4:len(in)])) * prime64x1
		h = rotl64_23(h)*prime64x2 + prime64x3
	}

	for ; i < len(in); i++ {
		h ^= uint64(in[i]) * prime64x5
		h = rotl64_11(h) * prime64x1
	}

	h ^= h >> 33
	h *= prime64x2
	h ^= h >> 29
	h *= prime64x3
	h ^= h >> 32

	return h
}
