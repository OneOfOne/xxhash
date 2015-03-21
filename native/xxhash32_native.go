package xxhash

import "hash"

const (
	prime32x1 uint32 = 2654435761
	prime32x2        = 2246822519
	prime32x3        = 3266489917
	prime32x4        = 668265263
	prime32x5        = 374761393
)

func rotl32(x, r uint32) uint32 {
	return (x << r) | (x >> (32 - r))
}

// Checksum32S returns the checksum of the input bytes with the specific seed.
func Checksum32S(in []byte, seed uint32) (h uint32) {
	i, l := 0, len(in)
	br := newbyteReader(in)
	if l >= 16 {
		var (
			v1 = seed + prime32x1 + prime32x2
			v2 = seed + prime32x2
			v3 = seed + 0
			v4 = seed - prime32x1
		)
		for ; i < l-16; i += 16 {
			v1 += br.Uint32(i) * prime32x2
			v1 = rotl32(v1, 13)
			v1 *= prime32x1

			v2 += br.Uint32(i+4) * prime32x2
			v2 = rotl32(v2, 13)
			v2 *= prime32x1

			v3 += br.Uint32(i+8) * prime32x2
			v3 = rotl32(v3, 13)
			v3 *= prime32x1

			v4 += br.Uint32(i+12) * prime32x2
			v4 = rotl32(v4, 13)
			v4 *= prime32x1
		}

		h = rotl32(v1, 1) + rotl32(v2, 7) + rotl32(v3, 12) + rotl32(v4, 18)

	} else {
		h = seed + prime32x5
	}

	h += uint32(l)
	for ; i < l-4; i += 4 {
		h += br.Uint32(i) * prime32x3
		h = rotl32(h, 17) * prime32x4
	}

	for ; i < l; i++ {
		h += uint32(br.Byte(i)) * prime32x5
		h = rotl32(h, 11) * prime32x1
	}

	h ^= h >> 15
	h *= prime32x2
	h ^= h >> 13
	h *= prime32x3
	h ^= h >> 16

	return
}

// Checksum32 returns the checksum of the input data with the seed set to 0
func Checksum32(in []byte) uint32 {
	return Checksum32S(in, 0)
}

type xxHash32 struct {
	ln                   uint64
	seed, v1, v2, v3, v4 uint32
	mem                  []byte
}

// Size returns the number of bytes Sum will return.
func (xx *xxHash32) Size() int {
	return 4
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (xx *xxHash32) BlockSize() int {
	return 16
}

// NewS32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting with the specific seed.
func NewS32(seed uint32) (xx hash.Hash32) {
	xx = &xxHash32{
		seed: seed,
		mem:  make([]byte, 16),
	}
	xx.Reset()
	return
}

// New32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting with the seed set to 0x0.
func New32() hash.Hash32 {
	return NewS32(0x0)
}

func (xx *xxHash32) Reset() {
	xx.v1 = xx.seed + prime32x1 + prime32x2
	xx.v2 = xx.seed + prime32x2
	xx.v3 = xx.seed
	xx.v4 = xx.seed - prime32x1
	xx.ln = 0
	xx.mem = xx.mem[:0]
}

func (xx *xxHash32) Write(in []byte) (n int, err error) {
	i, l, ml := 0, len(in), len(xx.mem)
	xx.ln += uint64(l)

	if d := 16 - ml; ml > 0 && ml+l > 16 {
		xx.mem = append(xx.mem, in[:d]...)
		in = in[d:]
		ml, l = 16, len(in)
	} else if ml+l <= 16 {
		xx.mem = append(xx.mem, in...)
		return l, nil
	}

	if ml > 0 {
		i += 16 - ml
		br := newbyteReader(xx.mem)
		xx.mem = append(xx.mem, in[:i]...)

		xx.v1 += br.Uint32(i) * prime32x2
		xx.v1 = rotl32(xx.v1, 13)
		xx.v1 *= prime32x1

		xx.v2 += br.Uint32(i+4) * prime32x2
		xx.v2 = rotl32(xx.v2, 13)
		xx.v2 *= prime32x1

		xx.v3 += br.Uint32(i+8) * prime32x2
		xx.v3 = rotl32(xx.v3, 13)
		xx.v3 *= prime32x1

		xx.v4 += br.Uint32(i+12) * prime32x2
		xx.v4 = rotl32(xx.v4, 13)
		xx.v4 *= prime32x1

		xx.mem = xx.mem[:0]
	}
	br := newbyteReader(in)
	if l >= 16 {
		for ; i < l-16; i += 16 {
			xx.v1 += br.Uint32(i) * prime32x2
			xx.v1 = rotl32(xx.v1, 13)
			xx.v1 *= prime32x1

			xx.v2 += br.Uint32(i+4) * prime32x2
			xx.v2 = rotl32(xx.v2, 13)
			xx.v2 *= prime32x1

			xx.v3 += br.Uint32(i+8) * prime32x2
			xx.v3 = rotl32(xx.v3, 13)
			xx.v3 *= prime32x1

			xx.v4 += br.Uint32(i+12) * prime32x2
			xx.v4 = rotl32(xx.v4, 13)
			xx.v4 *= prime32x1
		}

	}

	if l-i != 0 {
		xx.mem = append(xx.mem, in[i:]...)
	}

	if l-i > 16 {
		panic("len(in) - i > 16")
	}

	if cap(xx.mem) > 16 {
		panic("cap(xx.mem) > 16")
	}

	return l, nil
}

func (xx *xxHash32) Sum32() (h uint32) {
	i, l := 0, len(xx.mem)
	if xx.ln >= 16 {
		h = rotl32(xx.v1, 1) + rotl32(xx.v2, 7) + rotl32(xx.v3, 12) + rotl32(xx.v4, 18)
	} else {
		h = xx.seed + prime32x5
	}

	h += uint32(xx.ln)

	if len(xx.mem) > 0 {
		br := newbyteReader(xx.mem)
		for ; i < l-4; i += 4 {
			h += br.Uint32(i) * prime32x3
			h = rotl32(h, 17) * prime32x4
		}

		for ; i < l; i++ {
			h += uint32(br.Byte(i)) * prime32x5
			h = rotl32(h, 11) * prime32x1
		}
	}
	h ^= h >> 15
	h *= prime32x2
	h ^= h >> 13
	h *= prime32x3
	h ^= h >> 16

	return
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (xx *xxHash32) Sum(in []byte) []byte {
	s := xx.Sum32()
	return append(in, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}
