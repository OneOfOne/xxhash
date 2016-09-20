// +build appengine safe

package xxhash

import "io"

// Backend returns the current version of xxhash being used.
const Backend = "GoSafe"

func ChecksumString32S(s string, seed uint32) uint32 {
	return Checksum32S([]byte(s), seed)
}

func ChecksumString64S(s string, seed uint64) uint64 {
	return Checksum64S([]byte(s), seed)
}

func writeString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(s))
}
