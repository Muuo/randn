# Golang Random String Generator

randn.go contains several implementations of the algo. You can use randn_test.go to test and benchmark.

```bash
$ go test
PASS
ok      randn   0.531s
```
```bash
$ go test -bench=.
PASS
BenchmarkLookup-4        5000000               344 ns/op
BenchmarkBLookup-4       2000000               867 ns/op
BenchmarkALookup-4       5000000               280 ns/op
BenchmarkAALookup-4     20000000               134 ns/op
ok      randn   13.025s
```
