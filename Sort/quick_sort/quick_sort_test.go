package quick_sort

import (
	"Algorithm/test/data"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testData := data.RandomData(1000)
	QuickSort(testData)
}
