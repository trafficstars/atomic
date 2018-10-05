```
BenchmarkSimpleStringMapGet-4           30000000                40.4 ns/op             0 B/op          0 allocs/op
BenchmarkAtomicStringMapGet-4           30000000                44.8 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleUint64MapGet-4           100000000               20.5 ns/op             0 B/op          0 allocs/op
BenchmarkAtomicUint64MapGet-4           50000000                29.1 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleUint64MapSetMap-4          100000             23535 ns/op           34466 B/op        427 allocs/op
BenchmarkAtomicUint64MapSetMap-4           50000             24181 ns/op           35526 B/op        439 allocs/op
```
