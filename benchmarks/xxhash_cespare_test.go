package exttests

import (
	"testing"

	"github.com/cespare/xxhash"
)

const inS = `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`

var (
	in = []byte(inS)
)

func BenchmarkXXSum64Cespare(b *testing.B) {
	var bv uint64
	b.Run("Func", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bv += xxhash.Sum64(in)
		}
	})
	b.Run("Struct", func(b *testing.B) {
		h := xxhash.New()
		for i := 0; i < b.N; i++ {
			h.Write(in)
			bv += h.Sum64()
			h.Reset()
		}
	})
}

func BenchmarkXXSum64ShortCespare(b *testing.B) {
	var bv uint64
	k := []byte("Test-key-100")
	b.Run("Func", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bv += xxhash.Sum64(k)
		}
	})
	b.Run("Struct", func(b *testing.B) {
		h := xxhash.New()
		for i := 0; i < b.N; i++ {
			h.Write(k)
			bv += h.Sum64()
			h.Reset()
		}
	})
}
