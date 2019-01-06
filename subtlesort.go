// Package subtlesort provides primitives for sorting slices and user-defined
// collections in constant-time.
package subtlesort

// A type that satisfies subtlesort.Interface can be sorted in constant-time
// by the methods of subtlesort.
type Interface interface {
	// Len is the number of elements in the collection
	Len() int

	// MinMax checks whether the elements with index i and j are in the right
	// order and swaps them if they are not.
	//
	//
	// Warning: subtlesort.Sort is only constant-time if MinMax is constant-time
	MinMax(i, j int)
}

// Sort sorts data in constant-time (assuming data.MinMax is runs in
// constant-time).
func Sort(data Interface) {
	n := data.Len()
	if n == 1 {
		return
	}
	if n == 2 {
		data.MinMax(0, 1)
		return
	}
	if n == 3 {
		data.MinMax(0, 1)
		data.MinMax(1, 2)
		data.MinMax(0, 1)
		return
	}

	t := 1
	for t <= n {
		t *= 2
	}
	t >>= 1

	p := t
	for p > 0 {
		q := t
		r := 0
		d := p
		for d > 0 {
			for i := 0; i < n-d; i++ {
				if i&p == r {
					data.MinMax(i, d+i)
				}
			}
			d = q - p
			q >>= 1
			r = p
		}
		p >>= 1
	}
}

func Int64s(data []int64) {
	Sort(Int64Slice(data))
}

func Int32s(data []int32) {
	Sort(Int32Slice(data))
}

// Int64Slices attaches the methods of Interface to []int64, sorting in the
// usual order.
type Int64Slice []int64

func (p Int64Slice) Len() int {
	return len(p)
}
func (p Int64Slice) MinMax(i, j int) {
	a := p[i]
	b := p[j]
	ab := a ^ b
	c := b - a
	c ^= ab & (c ^ b)
	c >>= 31
	c &= ab
	p[i] ^= c
	p[j] ^= c
}

// Int32Slices attaches the methods of Interface to []int32, sorting in the
// usual order.
type Int32Slice []int32

func (p Int32Slice) Len() int {
	return len(p)
}
func (p Int32Slice) MinMax(i, j int) {
	a := p[i]
	b := p[j]
	ab := a ^ b
	c := b - a
	c ^= ab & (c ^ b)
	c >>= 31
	c &= ab
	p[i] ^= c
	p[j] ^= c
}
