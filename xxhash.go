package xxhash

import "errors"

var (
	// ErrAlreadyComputed is returned if you try to call Write after calling Sum
	ErrAlreadyComputed = errors.New("cannot update an already computed hash")
	// ErrMemoryLimit is returned if you try to call Write with more than 1 Gigabytes of data
	ErrMemoryLimit = errors.New("cannot use more than 1 Gigabytes of data at once")
)

const (
	oneGb = 1 << 30
)
