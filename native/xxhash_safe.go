// +build appengine

package xxhash

// Backend returns the current version of xxhash being used.
const Backend = "Go"

func readU32le(b []byte, i int) uint32 {
	b = b[i:]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func readU64le(b []byte, i int) uint64 {
	b = b[i:]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
