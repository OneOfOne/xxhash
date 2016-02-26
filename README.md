# xxhash [![GoDoc](http://godoc.org/github.com/OneOfOne/xxhash?status.svg)](http://godoc.org/github.com/OneOfOne/xxhash) [![Build Status](https://travis-ci.org/OneOfOne/xxhash.svg?branch=master)](https://travis-ci.org/OneOfOne/xxhash)
--

This is a CGO wrapper and a native Go implementation of the excellent [xxhash](https://github.com/Cyan4973/xxHash)* algorithm, an extremely fast non-cryptographic Hash algorithm, working at speeds close to RAM limits.

* The C implementation is ([Copyright](https://github.com/Cyan4973/xxHash/blob/master/LICENSE) (c) 2012-2014, Yann Collet)

## Install

Install *the recommended* pure-go version (much faster with shorter inputs):

	go get github.com/OneOfOne/xxhash/native

Or to install the CGO wrapper over the original C code (only recommended if hashing huge slices at a time):

	go get github.com/OneOfOne/xxhash

## Features

* The native version is optimized and is as fast as you can get in pure Go.
* The native version falls back to a less optimized version on appengine (it uses unsafe).
* Both the native version and the cgo version supports 64bit and 32bit versions of the algorithm.
* When using the cgo version, it will automatically fallback to the native version if cgo isn't available, you can check the `xxhash.Backend` const.

## Benchmark

	go test github.com/OneOfOne/xxhash -bench=. -benchmem

### Core i7-4790 @ 3.60GHz, Linux 4.4.1 (64bit), Go dev.ssa (+d3f15ff 2016-02-25)

	BenchmarkXXChecksum32-8                   5000000               499 ns/op
	BenchmarkXXChecksum32Cgo-8                5000000               601 ns/op

	BenchmarkXXChecksumString32-8             5000000               492 ns/op
	BenchmarkXXChecksumString32Cgo-8          5000000               607 ns/op

	BenchmarkXXChecksum64-8                  10000000               262 ns/op
	BenchmarkXXChecksum64Cgo-8               10000000               447 ns/op

	BenchmarkXXChecksumString64-8            10000000               274 ns/op
	BenchmarkXXChecksumString64Cgo-8         10000000               459 ns/op

	BenchmarkXXChecksum64Short-8            300000000              10.3 ns/op
	BenchmarkXXChecksum64ShortCgo-8          10000000               285 ns/op

	BenchmarkXXChecksumString64Short-8      200000000              12.2 ns/op
	BenchmarkXXChecksumString64CgoShort-8    10000000               295 ns/op

	BenchmarkXX64MultiWrites-8                5000000               490 ns/op
	BenchmarkXX64CgoMultiWrites-8             5000000               819 ns/op


### Core i7-4790 @ 3.60GHz, Linux 4.0 (64bit), Go 1.5 (+433af05 2015-04-30)

	BenchmarkXxhash32                2000000              1060 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash32Cgo             3000000               492 ns/op               0 B/op          0 allocs/op

	BenchmarkXxhash64                2000000               762 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash64Cgo             5000000               359 ns/op               0 B/op          0 allocs/op

	BenchmarkFnv32                    500000              2400 ns/op              16 B/op          1 allocs/op
	BenchmarkFnv64                    500000              2427 ns/op              16 B/op          1 allocs/op
	BenchmarkAdler32                 1000000              1223 ns/op               0 B/op          0 allocs/op
	BenchmarkCRC32IEEE                300000              5354 ns/op               0 B/op          0 allocs/op

	BenchmarkXxhash64VeryShort     100000000              18.6 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash64CgoVeryShort   10000000               168 ns/op               0 B/op          0 allocs/op

	BenchmarkFnv64VeryShort         20000000              98.4 ns/op              32 B/op          2 allocs/op

	BenchmarkXxhash64MultiWrites     2000000               724 ns/op               0 B/op          0 allocs/op
	BenchmarkXxhash64CgoMultiWrites  5000000               331 ns/op               0 B/op          0 allocs/op

	BenchmarkFnv64MultiWrites        1000000              2357 ns/op               0 B/op          0 allocs/op

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

Copyright 2015-2016 Ahmed <[OneOfOne](https://github.com/OneOfOne/)> W.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
