# Caching

```go
$ go test -run none -bench . -benchtime 3s
Elements in the link list 16777216
Elements in the matrix 16777216
goos: darwin
goarch: amd64
pkg: github.com/dtherhtun/Learning-go/UG/caching
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkLinkListTraverse-16    	     127	  27741402 ns/op
BenchmarkColumnTraverse-16      	      30	 119340750 ns/op
BenchmarkRowTraverse-16         	     296	  12244211 ns/op
PASS
ok  	github.com/dtherhtun/Learning-go/UG/caching	19.968s
```
