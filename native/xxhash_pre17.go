// +build !go1.7

package xxhash

// Checksum32S returns the checksum of the input bytes with the specific seed.
func Checksum32S(in []byte, seed uint32) (h uint32) {
	var i, l int32 = 0, int32(len(in))
	br := newbyteReader(in)
	if l >= 16 {
		var (
			v1 = seed + prime32x1 + prime32x2
			v2 = seed + prime32x2
			v3 = seed + 0
			v4 = seed - prime32x1
		)
		for ; i <= l-16; i += 16 {
			v1 += br.Uint32(i) * prime32x2
			v1 = rotl32_13(v1)
			v1 *= prime32x1

			v2 += br.Uint32(i+4) * prime32x2
			v2 = rotl32_13(v2)
			v2 *= prime32x1

			v3 += br.Uint32(i+8) * prime32x2
			v3 = rotl32_13(v3)
			v3 *= prime32x1

			v4 += br.Uint32(i+12) * prime32x2
			v4 = rotl32_13(v4)
			v4 *= prime32x1
		}

		h = rotl32_1(v1) + rotl32_7(v2) + rotl32_12(v3) + rotl32_18(v4)

	} else {
		h = seed + prime32x5
	}

	h += uint32(l)
	for ; i <= l-4; i += 4 {
		h += br.Uint32(i) * prime32x3
		h = rotl32_17(h) * prime32x4
	}

	for ; i < l; i++ {
		h += uint32(br.Byte(i)) * prime32x5
		h = rotl32_11(h) * prime32x1
	}

	h ^= h >> 15
	h *= prime32x2
	h ^= h >> 13
	h *= prime32x3
	h ^= h >> 16

	return
}

func (xx *XXHash32) Write(in []byte) (int, error) {
	var i, l, ml int32 = 0, int32(len(in)), xx.memIdx
	xx.ln += int32(l)

	if d := 16 - ml; ml > 0 && ml+l > 16 {
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in[:d]))
		in = in[d:]
		ml, l = 16, int32(len(in))
	} else if ml+l < 16 {
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in))
		return int(l), nil
	}

	if ml > 0 {
		i += 16 - ml
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in))
		br := newbyteReader(xx.mem[:16])

		xx.v1 += br.Uint32(i) * prime32x2
		xx.v1 = rotl32_13(xx.v1)
		xx.v1 *= prime32x1

		xx.v2 += br.Uint32(i+4) * prime32x2
		xx.v2 = rotl32_13(xx.v2)
		xx.v2 *= prime32x1

		xx.v3 += br.Uint32(i+8) * prime32x2
		xx.v3 = rotl32_13(xx.v3)
		xx.v3 *= prime32x1

		xx.v4 += br.Uint32(i+12) * prime32x2
		xx.v4 = rotl32_13(xx.v4)
		xx.v4 *= prime32x1

		xx.memIdx = 0
	}
	br := newbyteReader(in)
	if l >= 16 {
		for ; i <= l-16; i += 16 {
			xx.v1 += br.Uint32(i) * prime32x2
			xx.v1 = rotl32_13(xx.v1)
			xx.v1 *= prime32x1

			xx.v2 += br.Uint32(i+4) * prime32x2
			xx.v2 = rotl32_13(xx.v2)
			xx.v2 *= prime32x1

			xx.v3 += br.Uint32(i+8) * prime32x2
			xx.v3 = rotl32_13(xx.v3)
			xx.v3 *= prime32x1

			xx.v4 += br.Uint32(i+12) * prime32x2
			xx.v4 = rotl32_13(xx.v4)
			xx.v4 *= prime32x1
		}

	}

	if l-i != 0 {
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in[i:]))
	}

	if debug {
		if l-i > 16 {
			panic("len(in) - i > 16")
		}
	}

	return int(l), nil
}

func (xx *XXHash32) Sum32() (h uint32) {
	var i, l int32 = 0, xx.memIdx
	if xx.ln >= 16 {
		h = rotl32_1(xx.v1) + rotl32_7(xx.v2) + rotl32_12(xx.v3) + rotl32_18(xx.v4)
	} else {
		h = xx.seed + prime32x5
	}

	h += uint32(xx.ln)

	if xx.memIdx > 0 {
		br := newbyteReader(xx.mem[:xx.memIdx])
		for ; i <= l-4; i += 4 {
			h += br.Uint32(i) * prime32x3
			h = rotl32_17(h) * prime32x4
		}

		for ; i < l; i++ {
			h += uint32(br.Byte(i)) * prime32x5
			h = rotl32_11(h) * prime32x1
		}
	}
	h ^= h >> 15
	h *= prime32x2
	h ^= h >> 13
	h *= prime32x3
	h ^= h >> 16

	return
}

// Checksum64S returns the 64bit xxhash checksum for a single input
func Checksum64S(in []byte, seed uint64) (h uint64) {
	var i, l int32 = 0, int32(len(in))
	br := newbyteReader(in)
	if l >= 32 {
		var (
			v1 = seed + prime64x1 + prime64x2
			v2 = seed + prime64x2
			v3 = seed + 0
			v4 = seed - prime64x1
		)
		for ; i <= l-32; i += 32 {
			v1 += br.Uint64(i) * prime64x2
			v1 = rotl64_31(v1) * prime64x1

			v2 += br.Uint64(i+8) * prime64x2
			v2 = rotl64_31(v2) * prime64x1

			v3 += br.Uint64(i+16) * prime64x2
			v3 = rotl64_31(v3) * prime64x1

			v4 += br.Uint64(i+24) * prime64x2
			v4 = rotl64_31(v4) * prime64x1
		}

		h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)
		v1 *= prime64x2
		v1 = rotl64_31(v1)
		v1 *= prime64x1
		h ^= v1
		h = h*prime64x1 + prime64x4

		v2 *= prime64x2
		v2 = rotl64_31(v2)
		v2 *= prime64x1
		h ^= v2
		h = h*prime64x1 + prime64x4

		v3 *= prime64x2
		v3 = rotl64_31(v3)
		v3 *= prime64x1
		h ^= v3
		h = h*prime64x1 + prime64x4

		v4 *= prime64x2
		v4 = rotl64_31(v4)
		v4 *= prime64x1
		h ^= v4
		h = h*prime64x1 + prime64x4
	} else {
		h = seed + prime64x5
	}

	h += uint64(l)

	for ; i <= l-8; i += 8 {
		k := br.Uint64(i)
		k *= prime64x2
		k = rotl64_31(k)
		k *= prime64x1
		h ^= k
		h = rotl64_27(h)*prime64x1 + prime64x4
	}

	for ; i <= l-4; i += 4 {
		h ^= uint64(br.Uint32(i)) * prime64x1
		h = rotl64_23(h)*prime64x2 + prime64x3
	}

	for ; i < l; i++ {
		h ^= uint64(br.Byte(i)) * prime64x5
		h = rotl64_11(h) * prime64x1
	}

	h ^= h >> 33
	h *= prime64x2
	h ^= h >> 29
	h *= prime64x3
	h ^= h >> 32

	return h
}

func (xx *XXHash64) Write(in []byte) (n int, err error) {
	var i, l, ml int32 = 0, int32(len(in)), xx.memIdx
	xx.ln += int32(l)
	if d := 32 - ml; ml > 0 && ml+l > 32 {
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in[:d]))
		in = in[d:]
		ml, l = 32, int32(len(in))
	} else if ml+l < 32 {
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in))
		return int(l), nil
	}

	if ml > 0 {
		i += 32 - ml
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in))
		br := newbyteReader(xx.mem[:32])

		xx.v1 += br.Uint64(0) * prime64x2
		xx.v1 = rotl64_31(xx.v1)
		xx.v1 *= prime64x1

		xx.v2 += br.Uint64(8) * prime64x2
		xx.v2 = rotl64_31(xx.v2)
		xx.v2 *= prime64x1

		xx.v3 += br.Uint64(16) * prime64x2
		xx.v3 = rotl64_31(xx.v3)
		xx.v3 *= prime64x1

		xx.v4 += br.Uint64(24) * prime64x2
		xx.v4 = rotl64_31(xx.v4)
		xx.v4 *= prime64x1

		xx.memIdx = 0
	}
	br := newbyteReader(in)
	if l >= 32 {
		for ; i <= l-32; i += 32 {
			xx.v1 += br.Uint64(i) * prime64x2
			xx.v1 = rotl64_31(xx.v1)
			xx.v1 *= prime64x1

			xx.v2 += br.Uint64(i+8) * prime64x2
			xx.v2 = rotl64_31(xx.v2)
			xx.v2 *= prime64x1

			xx.v3 += br.Uint64(i+16) * prime64x2
			xx.v3 = rotl64_31(xx.v3)
			xx.v3 *= prime64x1

			xx.v4 += br.Uint64(i+24) * prime64x2
			xx.v4 = rotl64_31(xx.v4)
			xx.v4 *= prime64x1
		}

	}

	if l-i != 0 {
		xx.memIdx += int32(copy(xx.mem[xx.memIdx:], in[i:]))
	}

	if debug {
		if l-i > 32 {
			panic("len(in) - i > 32")
		}
	}

	return int(l), nil
}

func (xx *XXHash64) Sum64() (h uint64) {
	var i int32
	v1, v2, v3, v4 := xx.v1, xx.v2, xx.v3, xx.v4
	if xx.ln >= 32 {
		h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)

		v1 *= prime64x2
		v1 = rotl64_31(v1)
		v1 *= prime64x1
		h ^= v1
		h = h*prime64x1 + prime64x4

		v2 *= prime64x2
		v2 = rotl64_31(v2)
		v2 *= prime64x1
		h ^= v2
		h = h*prime64x1 + prime64x4

		v3 *= prime64x2
		v3 = rotl64_31(v3)
		v3 *= prime64x1
		h ^= v3
		h = h*prime64x1 + prime64x4

		v4 *= prime64x2
		v4 = rotl64_31(v4)
		v4 *= prime64x1
		h ^= v4
		h = h*prime64x1 + prime64x4
	} else {
		h = xx.seed + prime64x5
	}

	h += uint64(xx.ln)
	if xx.memIdx > 0 {
		br := newbyteReader(xx.mem[:xx.memIdx])
		for ; i <= xx.memIdx-8; i += 8 {
			k := br.Uint64(i)
			k *= prime64x2
			k = rotl64_31(k)
			k *= prime64x1
			h ^= k
			h = rotl64_27(h)*prime64x1 + prime64x4
		}

		for ; i <= xx.memIdx-4; i += 4 {
			h ^= uint64(br.Uint32(i)) * prime64x1
			h = rotl64_23(h)*prime64x2 + prime64x3
		}

		for ; i < xx.memIdx; i++ {
			h ^= uint64(br.Byte(i)) * prime64x5
			h = rotl64_11(h) * prime64x1
		}
		xx.memIdx = 0
	}
	h ^= h >> 33
	h *= prime64x2
	h ^= h >> 29
	h *= prime64x3
	h ^= h >> 32

	return
}
