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
	br = br[i : i+4]
	return uint32(br[0]) | uint32(br[1])<<8 | uint32(br[2])<<16 | uint32(br[3])<<24
}

func (br byteReader) Uint64(i int) uint64 {
	br = br[i : i+8]
	return uint64(br[0]) | uint64(br[1])<<8 | uint64(br[2])<<16 | uint64(br[3])<<24 |
		uint64(br[4])<<32 | uint64(br[5])<<40 | uint64(br[6])<<48 | uint64(br[7])<<56
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
