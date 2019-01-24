go-subtlesort
=============

Constant-time sorting for Go with an interface similar to the standard library
[sort package](https://golang.org/pkg/sort/).

[Documentation is on godoc.](https://godoc.org/github.com/bwesterb/go-subtlesort)

Currently, performance is 50% slower than buildin sort (for `[]int32`),
but there is room for massive improvements.

```
BenchmarkSubtleSortInt32s1-4    	30000000	        55.2 ns/op	 144.81 MB/s
BenchmarkSubtleSortInt32s2-4    	20000000	       102 ns/op	 155.85 MB/s
BenchmarkSubtleSortInt32s3-4    	10000000	       226 ns/op	 141.17 MB/s
BenchmarkSubtleSortInt32s4-4    	 2000000	       611 ns/op	 104.67 MB/s
BenchmarkSubtleSortInt32s5-4    	 1000000	      1781 ns/op	  71.84 MB/s
BenchmarkSubtleSortInt32s6-4    	  300000	      4999 ns/op	  51.20 MB/s
BenchmarkSubtleSortInt32s7-4    	  100000	     13298 ns/op	  38.50 MB/s
BenchmarkSubtleSortInt32s8-4    	   50000	     33955 ns/op	  30.16 MB/s
BenchmarkSubtleSortInt32s9-4    	   20000	     85390 ns/op	  23.98 MB/s
BenchmarkSubtleSortInt32s10-4   	   10000	    211829 ns/op	  19.34 MB/s
BenchmarkSubtleSortInt32s11-4   	    3000	    564152 ns/op	  14.52 MB/s
BenchmarkSubtleSortInt32s12-4   	    1000	   1271305 ns/op	  12.89 MB/s
BenchmarkSubtleSortInt32s13-4   	     500	   2876991 ns/op	  11.39 MB/s
BenchmarkSortInt32s1-4          	20000000	       113 ns/op	  70.53 MB/s
BenchmarkSortInt32s2-4          	10000000	       148 ns/op	 107.91 MB/s
BenchmarkSortInt32s3-4          	 5000000	       307 ns/op	 104.21 MB/s
BenchmarkSortInt32s4-4          	 2000000	       753 ns/op	  84.92 MB/s
BenchmarkSortInt32s5-4          	 1000000	      1720 ns/op	  74.39 MB/s
BenchmarkSortInt32s6-4          	  300000	      5279 ns/op	  48.49 MB/s
BenchmarkSortInt32s7-4          	  100000	     12941 ns/op	  39.56 MB/s
BenchmarkSortInt32s8-4          	   50000	     30961 ns/op	  33.07 MB/s
BenchmarkSortInt32s9-4          	   20000	     71305 ns/op	  28.72 MB/s
BenchmarkSortInt32s10-4         	   10000	    160705 ns/op	  25.49 MB/s
BenchmarkSortInt32s11-4         	    5000	    360543 ns/op	  22.72 MB/s
BenchmarkSortInt32s12-4         	    2000	    826594 ns/op	  19.82 MB/s
BenchmarkSortInt32s13-4         	    1000	   1771341 ns/op	  18.50 MB/s
```

