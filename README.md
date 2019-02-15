# skyencoder
Code-generation based encoder for Skycoin

[![Build Status](https://travis-ci.com/skycoin/skyencoder.svg?branch=master)](https://travis-ci.com/skycoin/skyencoder)
[![GoDoc](https://godoc.org/github.com/skycoin/skyencoder?status.svg)](https://godoc.org/github.com/skycoin/skyencoder)

## Introduction

`skyencoder` generates a file with encode and decode methods for an encodable struct, using the [Skycoin binary encoding format](https://github.com/skycoin/skycoin/wiki/Skycoin-Binary-Encoding-Format).

For encodable non-struct types, you can wrap the non-struct type in a struct for the same result. A struct definition adds no overhead and does not change the encoding.

Skycoin's [`package encoder`](https://godoc.org/github.com/skycoin/skycoin/src/cipher/encoder) has a reflect-based encoder that can be used at runtime,
and supports any encodable type.

## Installation

```sh
go get github.com/skycoin/skyencoder/cmd/skyencoder
```

This installs `skyencoder` to `$GOPATH/bin`.  Make sure `$GOPATH/bin` is in
your shell environment's `$PATH` variable in order to invoke it in the shell.

## go:generate

To use `go generate` to generate the code, add a directive like this in the file where the struct is defined:

```go
// go:generate skyencoder -struct Foo
```

Then, use `go:generate` to generate it:

```sh
go generate github.com/foo/foo
```

## CLI Usage

```
» go run cmd/skyencoder/skyencoder.go --help
Usage of skyencoder:
	skyencoder [flags] -struct T [go import path e.g. github.com/skycoin/skycoin/src/coin]
	skyencoder [flags] -struct T files... # Must be a single package
Flags:
  -no-test
    	disable generating the _test.go file (test files require github.com/google/go-cmp/cmp and github.com/skycoin/encodertest)
  -output-file string
    	output file name; default <struct_name>_skyencoder.go
  -output-path string
    	output path; defaults to the package's path, or the file's containing folder
  -package string
    	package name for the output; if not provided, defaults to the struct's package
  -silent
    	disable all non-error log output
  -struct string
    	struct name, must be set
  -tags string
    	comma-separated list of build tags to apply
  -unexported
    	don't export generated methods (always true if the struct is not an exported type)
```

`skyencoder` generates a file with encode and decode methods for a struct, using the [Skycoin encoding format](github.com/skycoin/skycoin/wiki/encoder).

By default, the generated file is written to the same package as the source struct type.

If you wish to have the file written to a different location, use `-package` to control the name of the destination package,
`-output-path` to control the destination path, and `-output-file` to control the destination filename.

Build tags can be applied to the loaded package with `-tags`.

## CLI Examples

Generate code for struct `coin.SignedBlock` in `github.com/skycoin/skycoin/src/coin`:

```sh
go run cmd/skyencoder/skyencoder.go -struct SignedBlock github.com/skycoin/skycoin/src/coin
```

Generate code for struct `Foo` in `/tmp/foo/foo.go`:

```sh
go run cmd/skyencoder/skyencoder.go -struct Foo /tmp/foo/foo.go
```

*Note: absolute paths can only point to a Go file. If there are multiple Go files in that same path, all of them must be included.*

Generate code for struct `coin.SignedBlock` in `github.com/skycoin/skycoin/src/coin`, but sent to an external package:

```sh
go run cmd/skyencoder/skyencoder.go -struct SignedBlock -package foo -output-path /tmp/foo github.com/skycoin/skycoin/src/coin
```

*Note: do not use `-package` if the generated file is going to be in the same package as the struct*

## Generated tests

A file with tests is generated by default and can be disabled with `-no-test`.
This test file requires `github.com/google/go-cmp/cmp` and `github.com/google/go-cmp/cmp/cmpopts`.

Autogenerated tests will check that encoding and decoding succeeds and that the output matches the reflect-based `github.com/skycoin/skycoin/src/cipher/encoder`.

Notes:

* Autogenerated tests do not cover maxlen exceeded errors

## Benchmark results

Benchmarks compare the reflect-based `github.com/skycoin/skycoin/src/cipher/encoder` to the generated encoder.
Benchmarks performed on a Mid-2015 base model 15" Macbook Pro.

Comparison of `skyencoder` to other encoders is available at:

* https://github.com/gz-c/skycoin-serialization-benchmarks
* https://github.com/gz-c/gosercomp

```
» make bench
go test -benchmem -bench '.*' ./benchmark
goos: darwin
goarch: amd64
pkg: github.com/skycoin/skyencoder/benchmark
BenchmarkEncodeSize-8                    	200000000	         6.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkCipherEncodeSize-8              	 1000000	      1488 ns/op	     128 B/op	      16 allocs/op
BenchmarkEncodeToBuffer-8                	10000000	       133 ns/op	       0 B/op	       0 allocs/op
BenchmarkEncode-8                        	10000000	       189 ns/op	     112 B/op	       1 allocs/op
BenchmarkCipherEncode-8                  	  500000	      3103 ns/op	     400 B/op	      34 allocs/op
BenchmarkDecode-8                        	 5000000	       375 ns/op	     120 B/op	      10 allocs/op
BenchmarkCipherDecode-8                  	  500000	      2425 ns/op	     472 B/op	      29 allocs/op
BenchmarkEncodeSizeSignedBlock-8         	50000000	        25.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCipherEncodeSizeSignedBlock-8   	  200000	      8125 ns/op	     600 B/op	      75 allocs/op
BenchmarkEncodeSignedBlockToBuffer-8     	 3000000	       418 ns/op	       0 B/op	       0 allocs/op
BenchmarkEncodeSignedBlock-8             	 2000000	       721 ns/op	    1792 B/op	       1 allocs/op
BenchmarkCipherEncodeSignedBlock-8       	  100000	     19673 ns/op	    4080 B/op	     185 allocs/op
BenchmarkDecodeSignedBlock-8             	 1000000	      1002 ns/op	    1648 B/op	      10 allocs/op
BenchmarkCipherDecodeSignedBlock-8       	  100000	     13919 ns/op	    5448 B/op	     130 allocs/op
PASS
ok  	github.com/skycoin/skyencoder/benchmark	24.198s
```
