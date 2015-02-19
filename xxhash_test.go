package xxhash_test

import (
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"testing"

	C "github.com/OneOfOne/xxhash"
	N "github.com/OneOfOne/xxhash/native"
)

var (
	in = []byte(`Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`)
)

const (
	expected32 uint32 = 0x6101218F
	expected64 uint64 = 0xFFAE31BEBFED7652
)

func Test(t *testing.T) {
	t.Logf("CGO version's backend: %s", C.Backend)
	t.Logf("Native version's backend: %s", N.Backend)
}

func TestHash32(t *testing.T) {
	h := N.New32()
	h.Write(in)
	r := h.Sum32()
	if r != expected32 {
		t.Errorf("expected 0x%x, got 0x%x.", expected32, r)
	}
}

func TestHash32Cgo(t *testing.T) {
	h := C.New32()
	h.Write(in)
	r := h.Sum32()
	if r != expected32 {
		t.Errorf("expected 0x%x, got 0x%x.", expected32, r)
	}
}

func TestHash32Short(t *testing.T) {
	r := N.Checksum32(in)
	if r != expected32 {
		t.Errorf("expected 0x%x, got 0x%x.", expected32, r)
	}
}

func TestHash32CgoShort(t *testing.T) {
	r := C.Checksum32(in)
	if r != expected32 {
		t.Errorf("expected 0x%x, got 0x%x.", expected32, r)
	}
}

func TestHash64(t *testing.T) {
	h := N.New64()
	h.Write(in)
	r := h.Sum64()
	if r != expected64 {
		t.Errorf("expected 0x%x, got 0x%x.", expected64, r)
	}
}

func TestHash64Cgo(t *testing.T) {
	h := C.New64()
	h.Write(in)
	r := h.Sum64()
	if r != expected64 {
		t.Errorf("expected 0x%x, got 0x%x.", expected64, r)
	}
}

func TestHash64Short(t *testing.T) {
	r := N.Checksum64(in)
	if r != expected64 {
		t.Errorf("expected 0x%x, got 0x%x.", expected64, r)
	}
}

func TestHash64CgoShort(t *testing.T) {
	r := C.Checksum64(in)
	if r != expected64 {
		t.Errorf("expected 0x%x, got 0x%x.", expected64, r)
	}
}

func BenchmarkXxhash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		N.Checksum32(in)
	}
}

func BenchmarkXxhash32Cgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C.Checksum32(in)
	}
}

func BenchmarkXxhash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		N.Checksum64(in)
	}
}

func BenchmarkXxhash64Cgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C.Checksum64(in)
	}
}

func BenchmarkFnv32(b *testing.B) {
	h := fnv.New32()
	for i := 0; i < b.N; i++ {
		h.Write(in)
		h.Sum(nil)
	}
}

func BenchmarkFnv64(b *testing.B) {
	h := fnv.New64()
	for i := 0; i < b.N; i++ {
		h.Write(in)
		h.Sum(nil)
	}
}

func BenchmarkAdler32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		adler32.Checksum(in)
	}
}

func BenchmarkCRC32IEEE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE(in)
	}
}

func BenchmarkXxhash64VeryShort(b *testing.B) {
	k := []byte("Test-key-100")
	for i := 0; i < b.N; i++ {
		N.Checksum64(k)
	}
}

func BenchmarkFnv64VeryShort(b *testing.B) {
	k := []byte("Test-key-100")
	for i := 0; i < b.N; i++ {
		h := fnv.New64()
		h.Write(k)
		h.Sum(nil)
	}
}

func BenchmarkXxhash64CgoVeryShort(b *testing.B) {
	k := []byte("Test-key-100")
	for i := 0; i < b.N; i++ {
		C.Checksum64(k)
	}
}

func BenchmarkXxhash64MultiWrites(b *testing.B) {
	h := N.New64()
	for i := 0; i < b.N; i++ {
		h.Write(in)
	}
	_ = h.Sum64()
}

func BenchmarkFnv64MultiWrites(b *testing.B) {
	h := fnv.New64()
	for i := 0; i < b.N; i++ {
		h.Write(in)
	}
	_ = h.Sum64()
}

func BenchmarkXxhash64CgoMultiWrites(b *testing.B) {
	h := C.New64()
	for i := 0; i < b.N; i++ {
		h.Write(in)
	}
	_ = h.Sum64()
}
