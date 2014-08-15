/*
[xxhash](https://code.google.com/p/xxhash/) (Copyright (c) 2012-2014, Yann Collet) is an extremely fast non-cryptographic Hash algorithm, working at speeds close to RAM limits.

Supports both the 32bit and 64bit versions of the the algorithm.


[![Build Status](https://travis-ci.org/OneOfOne/xxhash.svg?branch=master)](https://travis-ci.org/OneOfOne/xxhash)

## Install

Install *the recommended* pure-go version (much faster with shorter input):

	go get github.com/OneOfOne/xxhash/native

Or to install the CGO wrapper over the original C code (only recommended if hashing huge slices at a time):

	go get github.com/OneOfOne/xxhash

## Features

* The native version is optimized and almost as fast as you can get in pure Go.
* The native version falls back to a less optimized version on appengine (it uses unsafe).
* Both the native version and the cgo version supports 64bit and 32bit versions of the algorithm.
* When using the cgo version, it will automaticly fallback to the native version if cgo isn't available.

## License

[Apache v2.0](http://opensource.org/licenses/Apache-2.0)

Copyright [2014] [[OneOfOne](https://github.com/OneOfOne/)]

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package xxhash

import "fmt"

var (
	// ErrAlreadyComputed is returned if you try to call Write after calling Sum
	ErrAlreadyComputed = fmt.Errorf("cannot update an already computed hash")
	// ErrMemoryLimit is returned if you try to call Write with more than 1 Gigabytes of data
	ErrMemoryLimit = fmt.Errorf("cannot use more than 1 Gigabytes of data at once")
)

const (
	oneGb = 2 << 30
)
