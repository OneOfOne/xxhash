# xxhash
--
    import "github.com/OneOfOne/xxhash"

Package xxhash is an up to date Go wrapper for
[xxhash](https://code.google.com/p/xxhash/), xxHash is an extremely fast
non-cryptographic Hash algorithm, working at speeds close to RAM limits.

Supports both the 32bit and 64bit versions of the the algorithm.

## Install

```sh github.com/OneOfOne/xxhash ```

## Benchmark

    go test github.com/OneOfOne/xxhash -bench=. -benchmem

Core i5-750 @ 2.67GHz, Linux 3.15.6 (64bit)

    BenchmarkXxhash64        3000000               472 ns/op               0 B/op          0 allocs/op
    BenchmarkXxhash32        2000000               725 ns/op               0 B/op          0 allocs/op
    BenchmarkFnv32           1000000              2004 ns/op            2816 B/op          1 allocs/op
    BenchmarkFnv64           1000000              2022 ns/op            2816 B/op          1 allocs/op
    BenchmarkAdler32          500000              2618 ns/op               0 B/op          0 allocs/op
    BenchmarkCRC32IEEE        200000              8659 ns/op               0 B/op          0 allocs/op

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

## Usage

```go
var (
	// ErrAlreadyComputed is returned if you try to call Write after calling Sum
	ErrAlreadyComputed = fmt.Errorf("cannot update an already computed hash")
	// ErrMemoryLimit is returned if you try to call Write with more than 1 Gigabytes of data
	ErrMemoryLimit = fmt.Errorf("cannot use more than 1 Gigabytes of data at once")
)
```

#### func  Checksum32

```go
func Checksum32(in []byte) uint32
```
Checksum32 returns the checksum of the input data with the seed set to 0

#### func  Checksum32S

```go
func Checksum32S(in []byte, seed uint32) uint32
```
Checksum32S returns the checksum of the input bytes with the specific seed.

#### func  Checksum64

```go
func Checksum64(in []byte) uint64
```
Checksum64 returns the checksum of the input data with the seed set to 0

#### func  Checksum64S

```go
func Checksum64S(in []byte, seed uint64) uint64
```
Checksum64S returns the checksum of the input bytes with the specific seed.

#### func  New32

```go
func New32() hash.Hash32
```
New32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting
with the seed set to 0x0.

#### func  New64

```go
func New64() hash.Hash64
```
New64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting
with the seed set to 0x0.

#### func  NewS32

```go
func NewS32(seed uint32) hash.Hash32
```
NewS32 creates a new hash.Hash32 computing the 32bit xxHash checksum starting
with the specific seed.

#### func  NewS64

```go
func NewS64(seed uint64) hash.Hash64
```
NewS64 creates a new hash.Hash64 computing the 64bit xxHash checksum starting
with the specific seed.
