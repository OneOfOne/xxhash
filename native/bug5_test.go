package xxhash

import "testing"

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
