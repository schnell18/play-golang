## Introduction

Benckmarking helps us to understand how a go program uses CPU, memory, disk
space and network bandwidth.

## Basic Benckmarking

### Write Benckmarking code

Implement benchmark code using `testing.B` similar to:

    func BenchmarkMemoryWaste(b *testing.B) {
        for i := 0; i < b.N; i++ {
            MemoryWaste(100)
        }
    }

As a convention, the function name should start with `Benckmark`.
The `b.N` controls the loop count, and can be customized using CLI switch.

You can also run benchmarking in parallel:

    func BenchmarkMemoryWasteParallel(b *testing.B) {
        b.RunParallel(func(pb *testing.PB){
            for pb.Next() {
                MemoryWaste(100)
            }
        })
    }


With go 1.24, [testing.B.Loop][1] is recommended way to write benchmarking code:

    func Benchmark(b *testing.B) {
      ... setup ...
      for b.Loop() {
        // optional timer control for in-loop setup/cleanup
        ... code to measure ...
      }
      ... cleanup ...
    }

This new form avoids dead code elimination, simplifying setup and cleanup
code.

### Run CPU Benckmarking

Run CPU benchmarking by using the `go test` command wit the `-bench` switch to
select benchmarking methods to execute. For example:

    go test ./... -bench=^Benchmark

This command runs all benchmarking methods with a `Benchmark` prefix.

### Run Memory Benckmarking

Run memory benchmarking by using the `go test` command wit the `-benchmem` switch
to select benchmarking methods to execute. For example:

    go test ./... -bench=^BenchmarkMemoryWaste -benchmem -count=10

This command runs all benchmarking methods with a `BenchmarkMemoryWaste` prefix
by printing memory allocation statistics.

## Compare benchmarking results

The `benchstat` compares the difference of two benchmarking runs. This tool
can be installed using:

    go install golang.org/x/perf/cmd/benchstat@latest

Then generates two benchmarking results by:

    go test ./... -bench='^BenchmarkMemoryWaste$' -benchmem -count=5 > unoptimized.bench
    go test ./... -bench='^BenchmarkMemoryWaste2$' -benchmem -count=5 > optimized.bench

Use benchstat to compare the difference:

    benchstat unoptimized.bench optimized.bench

[1]: https://go.dev/blog/testing-b-loop
