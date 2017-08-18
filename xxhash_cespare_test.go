// +build cespare

package xxhash_test

import (
	"testing"

	"github.com/cespare/xxhash"
)

func BenchmarkCespareXXChecksum64(b *testing.B) {
	var bv uint64
	for i := 0; i < b.N; i++ {
		bv += xxhash.Sum64(in)
	}
}

func BenchmarkCespareXXChecksum64Short(b *testing.B) {
	var bv uint64
	k := []byte("Test-key-100")
	for i := 0; i < b.N; i++ {
		bv += xxhash.Sum64(k)
	}
}
