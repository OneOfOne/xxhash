# xxhash [![GoDoc](http://godoc.org/github.com/OneOfOne/xxhash?status.svg)](http://godoc.org/github.com/OneOfOne/xxhash) [![Build Status](https://travis-ci.org/OneOfOne/xxhash.svg?branch=master)](https://travis-ci.org/OneOfOne/xxhash)
--

[xxhash](https://code.google.com/p/xxhash/) ([Copyright](https://code.google.com/p/xxhash/source/browse/trunk/LICENSE) (c) 2012-2014, Yann Collet) is an extremely fast non-cryptographic Hash algorithm, working at speeds close to RAM limits.

Supports both the 32bit and 64bit versions of the the algorithm.

## Install

Install *the recommended* pure-go version (much faster with shorter input):

	go get github.com/OneOfOne/xxhash/native

Or to install the CGO wrapper over the original C code (only recommended if hashing huge slices at a time):

	go get github.com/OneOfOne/xxhash

## Features

* The native version is optimized and is as you can get in pure Go.
* The native version falls back to a less optimized version on appengine (it uses unsafe).
* The native version (on non-appengine builds) by default **ignores endianness**, if you require endianness safety (aka same hash on big and little endian systems), build with `-tags be`.
* Both the native version and the cgo version supports 64bit and 32bit versions of the algorithm.
* When using the cgo version, it will automaticly fallback to the native version if cgo isn't available.

## Benchmark

	go test github.com/OneOfOne/xxhash -bench=. -benchmem

### Core i7-4790 @ 3.60GHz, Linux 3.16.2 (64bit)

	BenchmarkXxhash32                1000000              1029 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash32Cgo             3000000               456 ns/op               0 B/op          0 allocs/op

	BenchmarkXxhash64                2000000               727 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash64Cgo             5000000               328 ns/op               0 B/op          0 allocs/op

	BenchmarkFnv32                   300000               4930 ns/op            2816 B/op          1 allocs/op
	BenchmarkFnv64                   300000               4994 ns/op            2816 B/op          1 allocs/op
	BenchmarkAdler32                 1000000              1238 ns/op               0 B/op          0 allocs/op
	BenchmarkCRC32IEEE               300000               5087 ns/op               0 B/op          0 allocs/op

	# key = []byte("Test-key-100")
	BenchmarkXxhash64VeryShort       100000000            16.8 ns/op             0 B/op          0 allocs/op
	BenchmarkXxhash64CgoVeryShort    10000000              174 ns/op               0 B/op          0 allocs/op

	# MultiWrites uses h.Write multiple times
	BenchmarkXxhash64MultiWrites     2000000               815 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash64CgoMultiWrites  5000000               342 ns/op               0 B/op          0 allocs/op

## Usage
	h := xxhash.New64()
	// r, err := os.Open("......")
	// defer f.Close()
	r := strings.NewReader(F)
	io.Copy(h, r)
	fmt.Println("xxhash.Backend:", xxhash.Backend)
	fmt.Println("File checksum:", h.Sum64())

[<kbd>playground</kbd>](http://play.golang.org/p/rhRN3RdQyd)

## License

[Apache v2.0](http://opensource.org/licenses/Apache-2.0)

Copyright 2014 [OneOfOne](https://github.com/OneOfOne/)

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
