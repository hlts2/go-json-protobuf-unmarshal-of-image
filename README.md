# Golang Unmarshal Benchmark
Benchmark of unmarshalling image data by JSON and Protocol Buffers.

## Description
When handling image data via API, there is a method to put it in JSON and Protocol Buffers as a relatively easy method. Therefore, I will examine the performance of the two methods.

## Installation

```
go get github.com/hlts2/go-json-protobuf-unmarshal-of-image
```

## Motivation
In golang, compare them to make a performance-oriented choice when handle image data with JSON or Protocol Buffers.

## Usage

### Image Data
You must ceate image files into the `go-json-protobuf-unmarshal-of-image/images`.

### Benchmark

```
go test -bench . -benchmem
```

## Benchmark Result
In my environment, the image size per sheet is about 20 KB.

The following is the result of unmarshaling multiple images.

```
goos: darwin
goarch: amd64
pkg: github.com/hlts2/go-json-protobuf-unmarshal-of-image
BenchmarkUnmarshal/JsonUnmarshalOf4files-4         	    2000	   1113538 ns/op	  122106 B/op	      19 allocs/op
--- BENCH: BenchmarkUnmarshal/JsonUnmarshalOf4files-4
    bench_test.go:130: protobuf size: 120040
    bench_test.go:130: protobuf size: 120040
    bench_test.go:130: protobuf size: 120040
BenchmarkUnmarshal/ProtoUnmarshalOf4files-4        	  100000	     20588 ns/op	   94072 B/op	      16 allocs/op
--- BENCH: BenchmarkUnmarshal/ProtoUnmarshalOf4files-4
    bench_test.go:137: protobuf size: 90026
    bench_test.go:137: protobuf size: 90026
    bench_test.go:137: protobuf size: 90026
    bench_test.go:137: protobuf size: 90026
BenchmarkUnmarshal/JsonUnmarshalOf10files-4        	     500	   2232193 ns/op	  256044 B/op	      43 allocs/op
--- BENCH: BenchmarkUnmarshal/JsonUnmarshalOf10files-4
    bench_test.go:130: protobuf size: 240012
    bench_test.go:130: protobuf size: 240012
    bench_test.go:130: protobuf size: 240012
BenchmarkUnmarshal/ProtoUnmarshalOf10files-4       	   30000	     40422 ns/op	  186296 B/op	      36 allocs/op
--- BENCH: BenchmarkUnmarshal/ProtoUnmarshalOf10files-4
    bench_test.go:137: protobuf size: 180002
    bench_test.go:137: protobuf size: 180002
    bench_test.go:137: protobuf size: 180002
    bench_test.go:137: protobuf size: 180002
BenchmarkUnmarshal/JsonUnmarshalOf20files-4        	     300	   4055124 ns/op	  462350 B/op	      77 allocs/op
--- BENCH: BenchmarkUnmarshal/JsonUnmarshalOf20files-4
    bench_test.go:130: protobuf size: 437960
    bench_test.go:130: protobuf size: 437960
    bench_test.go:130: protobuf size: 437960
BenchmarkUnmarshal/ProtoUnmarshalOf20files-4       	   20000	     81857 ns/op	  347192 B/op	      67 allocs/op
--- BENCH: BenchmarkUnmarshal/ProtoUnmarshalOf20files-4
    bench_test.go:137: protobuf size: 328462
    bench_test.go:137: protobuf size: 328462
    bench_test.go:137: protobuf size: 328462
    bench_test.go:137: protobuf size: 328462
BenchmarkUnmarshal/JsonUnmarshalOf50files-4        	     100	  11473404 ns/op	 1326843 B/op	     165 allocs/op
--- BENCH: BenchmarkUnmarshal/JsonUnmarshalOf50files-4
    bench_test.go:130: protobuf size: 1244700
    bench_test.go:130: protobuf size: 1244700
BenchmarkUnmarshal/ProtoUnmarshalOf50files-4       	    5000	    248897 ns/op	  993464 B/op	     152 allocs/op
--- BENCH: BenchmarkUnmarshal/ProtoUnmarshalOf50files-4
    bench_test.go:137: protobuf size: 933520
    bench_test.go:137: protobuf size: 933520
    bench_test.go:137: protobuf size: 933520
BenchmarkUnmarshal/JsonUnmarshalOf100files-4       	      50	  23937355 ns/op	 2751468 B/op	     299 allocs/op
--- BENCH: BenchmarkUnmarshal/JsonUnmarshalOf100files-4
    bench_test.go:130: protobuf size: 2589844
    bench_test.go:130: protobuf size: 2589844
BenchmarkUnmarshal/ProtoUnmarshalOf100files-4      	    2000	    553408 ns/op	 2054328 B/op	     285 allocs/op
--- BENCH: BenchmarkUnmarshal/ProtoUnmarshalOf100files-4
    bench_test.go:137: protobuf size: 1942376
    bench_test.go:137: protobuf size: 1942376
    bench_test.go:137: protobuf size: 1942376
BenchmarkUnmarshal/JsonUnmarshalOf150files-4       	      50	  36947225 ns/op	 4000284 B/op	     436 allocs/op
--- BENCH: BenchmarkUnmarshal/JsonUnmarshalOf150files-4
    bench_test.go:130: protobuf size: 3777716
    bench_test.go:130: protobuf size: 3777716
BenchmarkUnmarshal/ProtoUnmarshalOf150files-4      	    2000	    789530 ns/op	 2998072 B/op	     421 allocs/op
--- BENCH: BenchmarkUnmarshal/ProtoUnmarshalOf150files-4
    bench_test.go:137: protobuf size: 2833301
    bench_test.go:137: protobuf size: 2833301
    bench_test.go:137: protobuf size: 2833301

```

# Conclusion
The data amount of Protocol Buffers is smaller than JSON, and unmarshal can be performed at high speed.
