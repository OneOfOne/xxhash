package xxhash

import (
	"errors"
	"hash"
)

const (
	prime64x1 uint64 = 11400714785074694791
	prime64x2        = 14029467366897019727
	prime64x3        = 1609587929392839161
	prime64x4        = 9650029242287828579
	prime64x5        = 2870177450012600261
)

var (
	LenErr = errors.New("len(in) - i > 32")
	CapErr = errors.New("cap(xx.mem) > 32")
)

func rotl64(x, r uint64) uint64 {
	return (x << r) | (x >> (64 - r))
}

// Checksum64S returns the 64bit xxhash checksum for a single input
func Checksum64S(in []byte, seed uint64) (h uint64) {
	i, l := 0, len(in)
	br := newbyteReader(in)
	if l >= 32 {
		var (
			v1 = seed + prime64x1 + prime64x2
			v2 = seed + prime64x2
			v3 = seed + 0
			v4 = seed - prime64x1
		)
		for ; i < l-32; i += 32 {
			v1 += br.Uint64(i) * prime64x2
			v1 = rotl64(v1, 31) * prime64x1

			v2 += br.Uint64(i+8) * prime64x2
			v2 = rotl64(v2, 31) * prime64x1

			v3 += br.Uint64(i+16) * prime64x2
			v3 = rotl64(v3, 31) * prime64x1

			v4 += br.Uint64(i+24) * prime64x2
			v4 = rotl64(v4, 31) * prime64x1
		}

		h = rotl64(v1, 1) + rotl64(v2, 7) + rotl64(v3, 12) + rotl64(v4, 18)
		v1 *= prime64x2
		v1 = rotl64(v1, 31)
		v1 *= prime64x1
		h ^= v1
		h = h*prime64x1 + prime64x4

		v2 *= prime64x2
		v2 = rotl64(v2, 31)
		v2 *= prime64x1
		h ^= v2
		h = h*prime64x1 + prime64x4

		v3 *= prime64x2
		v3 = rotl64(v3, 31)
		v3 *= prime64x1
		h ^= v3
		h = h*prime64x1 + prime64x4

		v4 *= prime64x2
		v4 = rotl64(v4, 31)
		v4 *= prime64x1
		h ^= v4
		h = h*prime64x1 + prime64x4
	} else {
		h = seed + prime64x5
	}

	h += uint64(l)

	for ; i < l-8; i += 8 {
		k := br.Uint64(i)
		k *= prime64x2
		k = rotl64(k, 31)
		k *= prime64x1
		h ^= k
		h = rotl64(h, 27)*prime64x1 + prime64x4
	}

	for ; i < l-4; i += 4 {
		h ^= uint64(br.Uint32(i)) * prime64x1
		h = rotl64(h, 23)*prime64x2 + prime64x3
	}

	for ; i < l; i++ {
		h ^= uint64(br.Byte(i)) * prime64x5
		h = rotl64(h, 11) * prime64x1
	}

	h ^= h >> 33
	h *= prime64x2
	h ^= h >> 29
	h *= prime64x3
	h ^= h >> 32

	return h
}

// Checksum64 an alias for Checksum64S(in, 0)
func Checksum64(in []byte) uint64 {
	return Checksum64S(in, 0)
}

type xxHash64 struct {
	ln                   uint64
	seed, v1, v2, v3, v4 uint64
	mem                  []byte
}

// Size returns the number of bytes Sum will return.
func (xx *xxHash64) Size() int {
	return 8
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (xx *xxHash64) BlockSize() int {
	return 32
}

// NewS64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the specific seed.
func NewS64(seed uint64) (xx hash.Hash64) {
	xx = &xxHash64{
		seed: seed,
		mem:  make([]byte, 32),
	}
	xx.Reset()
	return
}

// New64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting with the seed set to 0x0.
func New64() hash.Hash64 {
	return NewS64(0)
}

func (xx *xxHash64) Reset() {
	xx.v1 = xx.seed + prime64x1 + prime64x2
	xx.v2 = xx.seed + prime64x2
	xx.v3 = xx.seed
	xx.v4 = xx.seed - prime64x1
	xx.ln = 0
	xx.mem = xx.mem[:0]
}

func (xx *xxHash64) Write(in []byte) (n int, err error) {
	i, l, ml := 0, len(in), len(xx.mem)
	xx.ln += uint64(l)
	if d := 32 - ml; ml > 0 && ml+l > 32 {
		xx.mem = append(xx.mem, in[:d]...)
		in = in[d:]
		ml, l = 32, len(in)
	} else if ml+l <= 32 {
		xx.mem = append(xx.mem, in...)
		return l, nil
	}

	if ml > 0 {
		i += 32 - ml
		xx.mem = append(xx.mem, in[:i]...)
		br := newbyteReader(xx.mem)

		xx.v1 += br.Uint64(0) * prime64x2
		xx.v1 = rotl64(xx.v1, 31)
		xx.v1 *= prime64x1

		xx.v2 += br.Uint64(8) * prime64x2
		xx.v2 = rotl64(xx.v2, 31)
		xx.v2 *= prime64x1

		xx.v3 += br.Uint64(16) * prime64x2
		xx.v3 = rotl64(xx.v3, 31)
		xx.v3 *= prime64x1

		xx.v4 += br.Uint64(24) * prime64x2
		xx.v4 = rotl64(xx.v4, 31)
		xx.v4 *= prime64x1

		xx.mem = xx.mem[:0]
	}
	br := newbyteReader(in)
	if l >= 32 {
		for ; i < l-32; i += 32 {
			xx.v1 += br.Uint64(i) * prime64x2
			xx.v1 = rotl64(xx.v1, 31)
			xx.v1 *= prime64x1

			xx.v2 += br.Uint64(i+8) * prime64x2
			xx.v2 = rotl64(xx.v2, 31)
			xx.v2 *= prime64x1

			xx.v3 += br.Uint64(i+16) * prime64x2
			xx.v3 = rotl64(xx.v3, 31)
			xx.v3 *= prime64x1

			xx.v4 += br.Uint64(i+24) * prime64x2
			xx.v4 = rotl64(xx.v4, 31)
			xx.v4 *= prime64x1
		}

	}

	if l-i != 0 {
		xx.mem = append(xx.mem, in[i:]...)
	}

	if l-i > 32 {
		return 0, LenErr
	}

	if cap(xx.mem) > 32 {
		return 0, CapErr
	}
	return l, nil
}

func (xx *xxHash64) Sum64() (h uint64) {
	i, l := 0, len(xx.mem)
	v1, v2, v3, v4 := xx.v1, xx.v2, xx.v3, xx.v4
	if xx.ln >= 32 {
		h = rotl64(v1, 1) + rotl64(v2, 7) + rotl64(v3, 12) + rotl64(v4, 18)

		v1 *= prime64x2
		v1 = rotl64(v1, 31)
		v1 *= prime64x1
		h ^= v1
		h = h*prime64x1 + prime64x4

		v2 *= prime64x2
		v2 = rotl64(v2, 31)
		v2 *= prime64x1
		h ^= v2
		h = h*prime64x1 + prime64x4

		v3 *= prime64x2
		v3 = rotl64(v3, 31)
		v3 *= prime64x1
		h ^= v3
		h = h*prime64x1 + prime64x4

		v4 *= prime64x2
		v4 = rotl64(v4, 31)
		v4 *= prime64x1
		h ^= v4
		h = h*prime64x1 + prime64x4
	} else {
		h = xx.seed + prime64x5
	}

	h += xx.ln
	if len(xx.mem) > 0 {
		br := newbyteReader(xx.mem)
		for ; i < l-8; i += 8 {
			k := br.Uint64(i)
			k *= prime64x2
			k = rotl64(k, 31)
			k *= prime64x1
			h ^= k
			h = rotl64(h, 27)*prime64x1 + prime64x4
		}

		for ; i < l-4; i += 4 {
			h ^= uint64(br.Uint32(i)) * prime64x1
			h = rotl64(h, 23)*prime64x2 + prime64x3
		}

		for ; i < l; i++ {
			h ^= uint64(br.Byte(i)) * prime64x5
			h = rotl64(h, 11) * prime64x1
		}
	}
	h ^= h >> 33
	h *= prime64x2
	h ^= h >> 29
	h *= prime64x3
	h ^= h >> 32

	return
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (xx *xxHash64) Sum(in []byte) []byte {
	s := xx.Sum64()
	return append(in, byte(s>>56), byte(s>>48), byte(s>>40), byte(s>>32), byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}
