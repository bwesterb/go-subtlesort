package subtlesort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/bwesterb/go-subtlesort"
)

func TestSort(t *testing.T) {
	for i := 1; i < 100; i++ {
		testSort(t, i)
	}
	for i := 7; i < 20; i++ {
		testSort(t, 1<<uint(i))
	}
	testSort(t, 82549)
	testSort(t, 11909)
	testSort(t, 104711)
}

func testSort(t *testing.T, n int) {
	xs := make([]int32, n)
	for i := 0; i < n; i++ {
		xs[i] = int32(i)
	}
	rand.Shuffle(n, func(i, j int) { xs[i], xs[j] = xs[j], xs[i] })
	subtlesort.Int32s(xs)
	if !sort.SliceIsSorted(xs, func(i, j int) bool { return xs[i] <= xs[j] }) {
		t.Fatalf("Did not properly sort [0 ... %d], got: %v", n, xs)
	}
}

func TestMinMax32(t *testing.T) {
	xs := []int32{0, 1}
	subtlesort.Int32Slice(xs).MinMax(0, 1)
	if xs[0] != 0 || xs[1] != 1 {
		t.Fatal()
	}
	xs = []int32{1, 0}
	subtlesort.Int32Slice(xs).MinMax(0, 1)
	if xs[0] != 0 || xs[1] != 1 {
		t.Fatal()
	}
}

func benchSubtleSortInt32s(b *testing.B, n int) {
	xs := make([]int32, 1<<uint(n))
	b.SetBytes(int64(len(xs) * 4))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		subtlesort.Int32s(xs)
	}
}

func BenchmarkSubtleSortInt32s1(b *testing.B)  { benchSubtleSortInt32s(b, 1) }
func BenchmarkSubtleSortInt32s2(b *testing.B)  { benchSubtleSortInt32s(b, 2) }
func BenchmarkSubtleSortInt32s3(b *testing.B)  { benchSubtleSortInt32s(b, 3) }
func BenchmarkSubtleSortInt32s4(b *testing.B)  { benchSubtleSortInt32s(b, 4) }
func BenchmarkSubtleSortInt32s5(b *testing.B)  { benchSubtleSortInt32s(b, 5) }
func BenchmarkSubtleSortInt32s6(b *testing.B)  { benchSubtleSortInt32s(b, 6) }
func BenchmarkSubtleSortInt32s7(b *testing.B)  { benchSubtleSortInt32s(b, 7) }
func BenchmarkSubtleSortInt32s8(b *testing.B)  { benchSubtleSortInt32s(b, 8) }
func BenchmarkSubtleSortInt32s9(b *testing.B)  { benchSubtleSortInt32s(b, 9) }
func BenchmarkSubtleSortInt32s10(b *testing.B) { benchSubtleSortInt32s(b, 10) }
func BenchmarkSubtleSortInt32s11(b *testing.B) { benchSubtleSortInt32s(b, 11) }
func BenchmarkSubtleSortInt32s12(b *testing.B) { benchSubtleSortInt32s(b, 12) }
func BenchmarkSubtleSortInt32s13(b *testing.B) { benchSubtleSortInt32s(b, 13) }

func benchSortInt32s(b *testing.B, n int) {
	xs := make([]int32, 1<<uint(n))
	b.SetBytes(int64(len(xs) * 4))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(xs, func(i, j int) bool {
			return xs[i] <= xs[j]
		})
	}
}

func BenchmarkSortInt32s1(b *testing.B)  { benchSortInt32s(b, 1) }
func BenchmarkSortInt32s2(b *testing.B)  { benchSortInt32s(b, 2) }
func BenchmarkSortInt32s3(b *testing.B)  { benchSortInt32s(b, 3) }
func BenchmarkSortInt32s4(b *testing.B)  { benchSortInt32s(b, 4) }
func BenchmarkSortInt32s5(b *testing.B)  { benchSortInt32s(b, 5) }
func BenchmarkSortInt32s6(b *testing.B)  { benchSortInt32s(b, 6) }
func BenchmarkSortInt32s7(b *testing.B)  { benchSortInt32s(b, 7) }
func BenchmarkSortInt32s8(b *testing.B)  { benchSortInt32s(b, 8) }
func BenchmarkSortInt32s9(b *testing.B)  { benchSortInt32s(b, 9) }
func BenchmarkSortInt32s10(b *testing.B) { benchSortInt32s(b, 10) }
func BenchmarkSortInt32s11(b *testing.B) { benchSortInt32s(b, 11) }
func BenchmarkSortInt32s12(b *testing.B) { benchSortInt32s(b, 12) }
func BenchmarkSortInt32s13(b *testing.B) { benchSortInt32s(b, 13) }
