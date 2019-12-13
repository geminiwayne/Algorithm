package quick_sort

import (
	"math/rand"
)

/**
 * @QuickSort
 * quick sort the array
 * param: array of int
 * return: sorted array
 */

func QuickSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	left, right := 0, len(A)-1
	pivot := rand.Int() % len(A)
	A[pivot], A[right] = A[right], A[pivot]

	for i := range A {
		if A[i] < A[right] {
			A[left], A[i] = A[i], A[left]
			left++
		}
	}

	A[left], A[right] = A[right], A[left]
	QuickSort(A[:left])
	QuickSort(A[left+1:])
	return A
}
