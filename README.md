It's a project for benchmarks.

```
BenchmarkSimpleStringMapGet-4           50000000                41.0 ns/op             0 B/op          0 allocs/op
BenchmarkAtomicStringMapGet-4           30000000                46.3 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleUint64MapGet-4           100000000               21.3 ns/op             0 B/op          0 allocs/op
BenchmarkAtomicUint64MapGet-4           50000000                29.4 ns/op             0 B/op          0 allocs/op
BenchmarkAtomicValueUint64MapGet-4      50000000                27.7 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleUint64MapSetMap-4           50000             23991 ns/op           34468 B/op        427 allocs/op
BenchmarkAtomicUint64MapSetMap-4           50000             26139 ns/op           35523 B/op        439 allocs/op
```
