# Systems performance

Slides:

- [Go performance tips](https://go-talks.appspot.com/github.com/marselester/perf/go-tips.slide)

Use `present` tool to view the slides:

```sh
$ git clone https://github.com/marselester/perf.git
$ cd ./perf/
$ go run golang.org/x/tools/cmd/present -notes
```

## Examples

```sh
# example/false-sharing
BenchmarkSumNaive-12                           6451    181812 ns/op
BenchmarkSumPadded-12                          9349    127138 ns/op
# example/loop-invariant
BenchmarkMultiplyNaive-12                      2545    454713 ns/op
BenchmarkMultiplyInvariant-12                  3445    342182 ns/op
# example/loop-tiling
BenchmarkAddNaive-12                            634    1758105 ns/op
BenchmarkAddTiling-12                           967    1168991 ns/op
# example/loop-unrolling
BenchmarkDotNaive-12                        1401841    857.5 ns/op
BenchmarkDotUnroll-12                       3035698    389.9 ns/op
BenchmarkDotBoundsCheckingElimination-12    4201716    281.9 ns/op
```
