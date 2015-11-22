# (Opinionated) Golang Random String Generator

Several implementations of random string generators. You can use randn_test.go to test and benchmark.

** Note: **
  - ALookup and AALookup are optimized for speed and generate a maximum of 10 characters. A larger number of characters will require modification of the code... which is conveniently left up to the reader.
  - The code uses golang's pseudo-random number generator, so no assurance can be given concerning the probability of collisions. I'll post some stats once I have the time to do proper investigation.

```
$ go test
PASS
ok      randn   0.144s
```
```
$ go test -bench=.
PASS
BenchmarkLookup-4        3000000               483 ns/op
BenchmarkBLookup-4       2000000               598 ns/op
BenchmarkALookup-4      10000000               174 ns/op
BenchmarkAALookup-4     20000000               109 ns/op
ok      randn   8.144s
```
