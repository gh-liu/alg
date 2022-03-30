package sort

import (
	"golang.org/x/exp/constraints"
)

// Insertion Implementation of insertion sort.
func Insertion[T constraints.Ordered](array []T) []T {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i
		for ; j > 0 && temp < array[j-1]; j-- {
			array[j] = array[j-1]
		}
		array[j] = temp
	}

	return array
}
