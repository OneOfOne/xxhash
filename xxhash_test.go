package xxhash

import (
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"testing"
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

func TestHash32(t *testing.T) {
	h := New32()
	h.Write(in)
	r := h.Sum32()
	if r != expected32 {
		t.Errorf("expected 0x%x, got 0x%x.", expected32, r)
	}
}

func TestHash64(t *testing.T) {
	h := New64()
	h.Write(in)
	r := h.Sum64()
	if r != expected64 {
		t.Errorf("expected 0x%x, got 0x%x.", expected64, r)
	}
}

func BenchmarkXxhash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Checksum64(in)
	}
}

func BenchmarkXxhash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Checksum32(in)
	}
}

func BenchmarkFnv32(b *testing.B) {
	h := fnv.New32()
	for i := 0; i < b.N; i++ {
		h.Sum(in)
	}
}

func BenchmarkFnv64(b *testing.B) {
	h := fnv.New64()
	for i := 0; i < b.N; i++ {
		h.Sum(in)
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
