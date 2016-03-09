// +build appengine safe

package xxhash

import "io"

// Backend returns the current version of xxhash being used.
const Backend = "GoSafe"

type byteReader []byte

func newbyteReader(in []byte) byteReader {
	return byteReader(in)
}

func (br byteReader) Uint32(i int) uint32 {
	return uint32(br[i]) | uint32(br[i+1])<<8 | uint32(br[i+2])<<16 | uint32(br[i+3])<<24
}

func (br byteReader) Uint64(i int) uint64 {
	return uint64(br[i]) | uint64(br[i+1])<<8 | uint64(br[i+2])<<16 | uint64(br[i+3])<<24 |
		uint64(br[i+4])<<32 | uint64(br[i+5])<<40 | uint64(br[i+6])<<48 | uint64(br[i+7])<<56
}

func (br byteReader) Byte(i int) byte {
	return br[i]
}

func ChecksumString32S(s string, seed uint32) uint32 {
	return Checksum32S([]byte(s), seed)
}

func ChecksumString64S(s string, seed uint64) uint64 {
	return Checksum64S([]byte(s), seed)
}

func writeString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(s))
}
