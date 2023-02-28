# geodb [![GoDoc](https://godoc.org/github.com/itsmontoya/geodb?status.svg)](https://godoc.org/github.com/itsmontoya/geodb) ![Status](https://img.shields.io/badge/status-beta-yellow.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/itsmontoya/geodb)](https://goreportcard.com/report/github.com/itsmontoya/geodb) ![Go Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

geodb library intended to look up if a lat/long falls within a radius or polygon geofence.

## Benchmarks
```
BenchmarkgeodbInserting-4       1000000        1484 ns/op         112 B/op         2 allocs/op
BenchmarkgeodbLookup1_000-4    10000000         193 ns/op           0 B/op         0 allocs/op
BenchmarkgeodbLookup10_000-4    1000000        1283 ns/op           0 B/op         0 allocs/op
BenchmarkgeodbLookup100_000-4      2000      733435 ns/op         240 B/op         4 allocs/op
BenchmarkgeodbLookup_USA-4         2000      892946 ns/op         240 B/op         4 allocs/op

BenchmarkRTreeInserting-4           30000       51169 ns/op       23305 B/op       835 allocs/op
BenchmarkRTreeLookup1_000-4          5000      320515 ns/op      174992 B/op      3293 allocs/op
BenchmarkRTreeLookup10_000-4          300     4874605 ns/op     2246816 B/op     32193 allocs/op
BenchmarkRTreeLookup100_000-4          50    44346148 ns/op    20695248 B/op    239878 allocs/op
BenchmarkRTreeLookup_USA-4             20    65078809 ns/op    30001074 B/op    335703 allocs/op
```

## Usage
``` go
package main

import (
	"fmt"

	"github.com/itsmontoya/geodb"
)

func main() {
	gt := New(1)
	gt.Insert("Portland", 45.5231, 122.6765, 1000 * 15) // 15km radius
	fmt.Println(gt.Matches(45.4312, 122.7715))
}
```
