# xxhash [![GoDoc](http://godoc.org/github.com/OneOfOne/xxhash?status.svg)](http://godoc.org/github.com/OneOfOne/xxhash) [![Build Status](https://travis-ci.org/OneOfOne/xxhash.svg?branch=master)](https://travis-ci.org/OneOfOne/xxhash)
--

This is a CGO wrapper and a native Go implementation of the excellent [xxhash](https://github.com/Cyan4973/xxHash)* algorithm, an extremely fast non-cryptographic Hash algorithm, working at speeds close to RAM limits.

* The C implementation is ([Copyright](https://github.com/Cyan4973/xxHash/blob/master/LICENSE) (c) 2012-2014, Yann Collet)

## Install

	go get github.com/OneOfOne/xxhash

## Features

* On Go 1.7+ the pure go version is faster than CGO for all inputs, on < 1.7 it is still faster for short inputs.
* Supports ChecksumString{32,64} xxhash{32,64}.WriteString, which uses no copies when it can, falls back to copy on appengine.
* The native version falls back to a less optimized version on appengine due to the lack of unsafe.
* Both the native version and the cgo version supports 64bit and 32bit versions of the algorithm.
* When using the cgo version, it will automatically fallback to the native version if cgo isn't available or on go 1.7,
you can check the `xxhash.Backend` const.

## Benchmark
```
# need the forcecgo flag to actually bench CGO on 1.7
go test github.com/OneOfOne/xxhash -bench=. -benchmem -tags forcecgo
```

### Core i7-4790 @ 3.60GHz, Linux 4.4.5 (64bit), Go tip (+fcd2a06 2016-03-28)

```bash
➜ go test -bench=. -benchtime 3s -tags forcecgo
BenchmarkXXChecksum32-8                         10000000               352 ns/op
BenchmarkXXChecksum32Cgo-8                      10000000               562 ns/op

BenchmarkXXChecksumString32-8                   10000000               361 ns/op
BenchmarkXXChecksumString32Cgo-8                10000000               572 ns/op

BenchmarkXXChecksum64-8                         20000000               188 ns/op
BenchmarkXXChecksum64Cgo-8                      10000000               418 ns/op

BenchmarkXXChecksumString64-8                   20000000               202 ns/op
BenchmarkXXChecksumString64Cgo-8                10000000               436 ns/op

BenchmarkFnv32-8                                 2000000              2438 ns/op
BenchmarkFnv64-8                                 2000000              2407 ns/op

BenchmarkAdler32-8                               3000000              1242 ns/op

BenchmarkCRC32IEEE-8                            20000000               206 ns/op
BenchmarkCRC32IEEEString-8                       5000000               715 ns/op

BenchmarkCRC64ISO-8                              1000000              4767 ns/op
BenchmarkCRC64ISOString-8                        1000000              5418 ns/op

BenchmarkXXChecksum64Short-8                    500000000                9.71 ns/op
BenchmarkXXChecksum64ShortCgo-8                 20000000               263 ns/op

BenchmarkXXChecksumString64Short-8              500000000               11.6 ns/op
BenchmarkXXChecksumString64CgoShort-8           20000000               271 ns/op

BenchmarkCRC32IEEEShort-8                       200000000               24.2 ns/op
BenchmarkCRC64ISOShort-8                        200000000               22.6 ns/op

BenchmarkFnv64Short-8                           100000000               58.7 ns/op

BenchmarkXX64MultiWrites-8                      20000000               231 ns/op
BenchmarkXX64CgoMultiWrites-8                    5000000               744 ns/op

BenchmarkFnv64MultiWrites-8                      2000000              2368 ns/op
PASS
ok      github.com/OneOfOne/xxhash      137.142s
```

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

This project is released under the Apache v2. licence. See [LICENCE](LICENCE) for more details.