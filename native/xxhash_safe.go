// +build appengine safe

package xxhash

import (
	"io"
	"strings"
)

// Backend returns the current version of xxhash being used.
const Backend = "GoSafe"

type byteReader []byte

func newbyteReader(in []byte) byteReader {
	return byteReader(in)
}

func (br byteReader) Uint32(i int) (u uint32) {
	u = uint32(br[i]) | uint32(br[i+1])<<8 | uint32(br[i+2])<<16 | uint32(br[i+3])<<24
	if IsBigEndian {
		u = swap32(u)
	}
}

func (br byteReader) Uint64(i int) (u uint64) {
	u = uint64(br[i]) | uint64(br[i+1])<<8 | uint64(br[i+2])<<16 | uint64(br[i+3])<<24 |
		uint64(br[i+4])<<32 | uint64(br[i+5])<<40 | uint64(br[i+6])<<48 | uint64(br[i+7])<<56
	if IsBigEndian {
		u = swap64(u)
	}
	return
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
	n, err := io.Copy(w, struct{ io.Reader }{strings.NewReader(s)})
	return int(n), err
}
