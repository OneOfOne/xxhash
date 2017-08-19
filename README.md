# xxhash [![GoDoc](http://godoc.org/github.com/OneOfOne/xxhash?status.svg)](http://godoc.org/github.com/OneOfOne/xxhash) [![Build Status](https://travis-ci.org/OneOfOne/xxhash.svg?branch=master)](https://travis-ci.org/OneOfOne/xxhash)

This is a native Go implementation of the excellent [xxhash](https://github.com/Cyan4973/xxHash)* algorithm, an extremely fast non-cryptographic Hash algorithm, working at speeds close to RAM limits.

* The C implementation is ([Copyright](https://github.com/Cyan4973/xxHash/blob/master/LICENSE) (c) 2012-2014, Yann Collet)

## Install

	go get github.com/OneOfOne/xxhash

## Features

* On Go 1.7+ the pure go version is faster than CGO for all inputs.
* Supports ChecksumString{32,64} xxhash{32,64}.WriteString, which uses no copies when it can, falls back to copy on appengine.
* The native version falls back to a less optimized version on appengine due to the lack of unsafe.
* Almost as fast as the mostly pure assembly version written by the briliant [cespare](https://github.com/cespare/xxhash), while also supporting seeds.

## Benchmark
### Core i7-4790 @ 3.60GHz, Linux 4.12.6-1-ARCH (64bit), Go tip (+ff90f4af66 2017-08-19)

```bash
N/A
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
