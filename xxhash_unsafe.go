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
	if len(s) == 0 {
		return Checksum32S(nil, seed)
	}
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum32S((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

func (xx *XXHash32) WriteString(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return xx.Write((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)])
}

// ChecksumString64S returns the checksum of the input data, without creating a copy, with the specific seed.
func ChecksumString64S(s string, seed uint64) uint64 {
	if len(s) == 0 {
		return Checksum64S(nil, seed)
	}

	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Checksum64S((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s):len(s)], seed)
}

func (xx *XXHash64) WriteString(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	ss := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return xx.Write((*[maxInt32]byte)(unsafe.Pointer(ss.Data))[:len(s)])
}

// Checksum64S returns the 64bit xxhash checksum for a single input
func checksum64S(in []byte, seed uint64) uint64 {
	var (
		wordsLen = int(uint32(len(in)) >> 3)
		words    = ((*[maxInt32]uint64)(unsafe.Pointer(&in[0])))[:wordsLen:wordsLen]

		h uint64 = prime64x5

		v1, v2, v3, v4 = resetVs64(seed)

		i int
	)

	if len(in) < 32 {
		goto LessThan32
	}

	for m := len(words) - 3; i < m; i += 4 {
		words := (*[4]uint64)(unsafe.Pointer(&words[i]))

		v1 = round64(v1, words[0])
		v2 = round64(v2, words[1])
		v3 = round64(v3, words[2])
		v4 = round64(v4, words[3])
	}

	h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)

	h = mergeRound64(h, v1)
	h = mergeRound64(h, v2)
	h = mergeRound64(h, v3)
	h = mergeRound64(h, v4)

LessThan32:
	h += uint64(len(in))

	for ; i < len(words); i++ {
		k := round64(0, words[i])
		h ^= k
		h = rotl64_27(h)*prime64x1 + prime64x4
	}

	if in = in[wordsLen<<3 : len(in):len(in)]; len(in) > 3 {
		words := (*[1]uint32)(unsafe.Pointer(&in[0]))
		h ^= uint64(words[0]) * prime64x1
		h = rotl64_23(h)*prime64x2 + prime64x3

		in = in[4:len(in):len(in)]
	}

	for _, b := range in {
		h ^= uint64(b) * prime64x5
		h = rotl64_11(h) * prime64x1
	}

	return mix64(h)
}

/*
func (x *xxh) Write(b []byte) (n int, err error) {
	n = len(b)
	x.total += len(b)

	if x.n+len(b) < 32 {
		// This new data doesn't even fill the current block.
		copy(x.mem[x.n:], b)
		x.n += len(b)
		return
	}

	if x.n > 0 {
		// Finish off the partial block.
		copy(x.mem[x.n:], b)
		x.v1 = round(x.v1, u64(x.mem[0:8]))
		x.v2 = round(x.v2, u64(x.mem[8:16]))
		x.v3 = round(x.v3, u64(x.mem[16:24]))
		x.v4 = round(x.v4, u64(x.mem[24:32]))
		b = b[32-x.n:]
		x.n = 0
	}

	if len(b) >= 32 {
		// One or more full blocks left.
		b = writeBlocks(x, b)
	}

	// Store any remaining partial block.
	copy(x.mem[:], b)
	x.n = len(b)

	return
}
*/

func (xx *XXHash64) Write(in []byte) (n int, err error) {
	xx.ln, n = xx.ln+uint64(len(in)), len(in)

	if int(xx.memIdx)+len(in) < 32 {
		xx.memIdx += int8(copy(xx.mem[xx.memIdx:len(xx.mem):len(xx.mem)], in))
		return
	}

	i, v1, v2, v3, v4 := 0, xx.v1, xx.v2, xx.v3, xx.v4
	if d := 32 - int(xx.memIdx); d > 0 && int(xx.memIdx)+len(in) > 31 {
		copy(xx.mem[xx.memIdx:len(xx.mem):len(xx.mem)], in[:d:len(in)])

		words := (*[4]uint64)(unsafe.Pointer(&xx.mem[0]))

		v1 = round64(v1, words[0])
		v2 = round64(v2, words[1])
		v3 = round64(v3, words[2])
		v4 = round64(v4, words[3])

		if in, xx.memIdx = in[d:len(in):len(in)], 0; len(in) == 0 {
			goto RET
		}
	}

	for ; i < len(in)-31; i += 32 {
		words := (*[4]uint64)(unsafe.Pointer(&in[i]))

		v1 = round64(v1, words[0])
		v2 = round64(v2, words[1])
		v3 = round64(v3, words[2])
		v4 = round64(v4, words[3])
	}

	if len(in)-i != 0 {
		xx.memIdx += int8(copy(xx.mem[xx.memIdx:len(xx.mem):len(xx.mem)], in[i:len(in):len(in)]))
	}
RET:
	xx.v1, xx.v2, xx.v3, xx.v4 = v1, v2, v3, v4

	return
}

func (xx *XXHash64) Sum64() (h uint64) {
	if xx.ln > 31 {
		v1, v2, v3, v4 := xx.v1, xx.v2, xx.v3, xx.v4
		h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)

		h = mergeRound64(h, v1)
		h = mergeRound64(h, v2)
		h = mergeRound64(h, v3)
		h = mergeRound64(h, v4)
	} else {
		h = xx.seed + prime64x5
	}

	h += uint64(xx.ln)
	if xx.memIdx > 0 {
		var (
			in       = xx.mem[:xx.memIdx:xx.memIdx]
			wordsLen = int(uint32(len(in)) >> 3)
			words    = ((*[maxInt32]uint64)(unsafe.Pointer(&in[0])))[:wordsLen:wordsLen]
		)

		for _, k := range words {
			h ^= round64(0, k)
			h = rotl64_27(h)*prime64x1 + prime64x4
		}

		if in = in[wordsLen<<3 : len(in):len(in)]; len(in) > 3 {
			words := (*[1]uint32)(unsafe.Pointer(&in[0]))

			h ^= uint64(words[0]) * prime64x1
			h = rotl64_23(h)*prime64x2 + prime64x3

			in = in[4:len(in):len(in)]
		}

		for _, b := range in {
			h ^= uint64(b) * prime64x5
			h = rotl64_11(h) * prime64x1
		}
	}

	return mix64(h)
}
