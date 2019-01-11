# Golang Unmarshal Benchmark
Benchmark of unmarshalling image data by JSON and Protocol Buffers

## Description


## Installation

```
go get github.com/hlts2/go-json-protobuf-unmarshal-of-image
```

## Motivation
In golang, compare them to make a performance-oriented choice when handle image data with JSON or Protocol Buffers

## Usage

### Image Data
You must ceate image files into the `go-json-protobuf-unmarshal-of-image/images`

### Benchmark

```
go test -bench . -benchmem
```

# Conclusion
