package xxhash

import (
	"bytes"
	"testing"

	cxhash "github.com/OneOfOne/xxhash"
)

func TestReset64(t *testing.T) {
	h := New64()

	//
	p1 := "http"
	p2 := "://"
	p3 := "www.marmiton.org"
	p4 := "/recettes/recherche.aspx"
	p5 := "?st=2&aqt=gateau&"

	url := p1 + p2 + p3 + p4 + p5

	// compute hash by parts
	h.Write([]byte(p1))
	h.Write([]byte(p2))
	h.Write([]byte(p3))
	h.Write([]byte(p4))
	h.Write([]byte(p5))
	s1 := h.Sum64()

	h.Reset()
	h.Write([]byte(url))
	s2 := h.Sum64()

	// should be the same, right ?
	if s1 != s2 {
		t.Errorf("s1 != s2 %x %x", s1, s2)
	}
}

func TestReset32(t *testing.T) {
	h := New32()

	//
	p1 := "http"
	p2 := "://"
	p3 := "www.marmiton.org"
	p4 := "/recettes/recherche.aspx"
	p5 := "?st=2&aqt=gateau&"

	url := p1 + p2 + p3 + p4 + p5

	// compute hash by parts
	h.Write([]byte(p1))
	h.Write([]byte(p2))
	h.Write([]byte(p3))
	h.Write([]byte(p4))
	h.Write([]byte(p5))
	s1 := h.Sum32()

	h.Reset()
	h.Write([]byte(url))
	s2 := h.Sum32()

	// should be the same, right ?
	if s1 != s2 {
		t.Errorf("s1 != s2 %x %x", s1, s2)
	}
}

// issue 8
func TestDataLen(t *testing.T) {
	for i := 4; i <= 8096; i += 4 {
		testEquality(t, bytes.Repeat([]byte("www."), i/4))
	}
}

func testEquality(t *testing.T, v []byte) {
	ch64, ch32 := cxhash.Checksum64(v), cxhash.Checksum32(v)

	if h := Checksum64(v); ch64 != h {
		t.Fatalf("Checksum64 doesn't match, len = %d, expected 0x%X, got 0x%X", len(v), ch64, h)
	}

	if h := Checksum32(v); ch32 != h {
		t.Fatalf("Checksum32 doesn't match, len = %d, expected 0x%X, got 0x%X", len(v), ch32, h)
	}

	h64 := New64()
	h64.Write(v)

	if h := h64.Sum64(); ch64 != h {
		t.Fatalf("Sum64() doesn't match, len = %d, expected 0x%X, got 0x%X", len(v), ch64, h)
	}

	h32 := New32()
	h32.Write(v)

	if h := h32.Sum32(); ch32 != h {
		t.Fatalf("Sum32() doesn't match, len = %d, expected 0x%X, got 0x%X", len(v), ch32, h)
	}
}
