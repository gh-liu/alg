package sort

import "golang.org/x/exp/constraints"

// Merge Implementation of merge sort.
func Merge[S ~[]E, E constraints.Ordered](items S) S {
	if len(items) < 2 {
		return items
	}

	mid := len(items) / 2
	left := Merge(items[:mid])
	right := Merge(items[mid:])

	return merge(left, right)
}

// MergeIter Iteration implementation of merge sort.
func MergeIter[S ~[]E, E constraints.Ordered](items S) S {
	length := len(items)
	for step := 1; step < length; step = 2 * step {
		for i := 0; i+step < length; i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:min(i+2*step, len(items))])
			copy(items[i:], tmp)
		}
	}
	return items
}

func merge[S ~[]E, E constraints.Ordered](l, r S) S {
	var final []E
	var i, j int
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			final = append(final, l[i])
			i++
		} else {
			final = append(final, r[j])
			j++
		}
	}
	for ; i < len(l); i++ {
		final = append(final, l[i])
	}
	for ; j < len(r); j++ {
		final = append(final, r[j])
	}
	return final
}

func min[T constraints.Ordered](values ...T) T {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
