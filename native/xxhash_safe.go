// +build appengine

package xxhash

// Backend returns the current version of xxhash being used.
const Backend = "Go"

type byteReader struct {
	p []byte
}

func newbyteReader(in *[]byte) *byteReader {
	return &byteReader{*in}
}

func (br *byteReader) Uint32(i int) (u uint32) {
	return uint32(br.p[i]) | uint32(br.p[i+1])<<8 | uint32(br.p[i+2])<<16 | uint32(br.p[i+3])<<24
}

func (br *byteReader) Uint64(i int) (u uint64) {
	return uint64(br.p[i]) | uint64(br.p[i+1])<<8 | uint64(br.p[i+2])<<16 | uint64(br.p[i+3])<<24 |
		uint64(br.p[i+4])<<32 | uint64(br.p[i+5])<<40 | uint64(br.p[i+6])<<48 | uint64(br.p[i+7])<<56
}

func (br *byteReader) Byte(i int) byte {
	return br.p[i]
}
