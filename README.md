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
You must ceate image files into the `go-json-protobuf-unmarshal-of-image/images`

### Benchmark

```
go test -bench . -benchmem
```

## Benchmark Result

# Conclusion
The data amount of Protocol Buffers is smaller than JSON, and unmarshal can be performed at high speed.
