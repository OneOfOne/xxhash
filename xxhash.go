/*
Package xxhash is an up to date Go wrapper for [xxhash](https://code.google.com/p/xxhash/),
xxHash is an extremely fast non-cryptographic Hash algorithm, working at speeds close to RAM limits.

Supports both the 32bit and 64bit versions of the the algorithm.

## Install

```sh
github.com/OneOfOne/xxhash
```

## Benchmark

```
go test github.com/OneOfOne/xxhash -bench=. -benchmem
```

Core i5-750 @ 2.67GHz, Linux 3.15.6 (64bit)

```
BenchmarkXxhash64        3000000               472 ns/op               0 B/op          0 allocs/op
BenchmarkXxhash32        2000000               725 ns/op               0 B/op          0 allocs/op
BenchmarkFnv32           1000000              2004 ns/op            2816 B/op          1 allocs/op
BenchmarkFnv64           1000000              2022 ns/op            2816 B/op          1 allocs/op
BenchmarkAdler32          500000              2618 ns/op               0 B/op          0 allocs/op
BenchmarkCRC32IEEE        200000              8659 ns/op               0 B/op          0 allocs/op
```

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

/*
//pull in xhash.c to compile it staticly
#cgo CFLAGS: -std=c99
#include "c-trunk/xxhash.c"
*/
import "C"
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
