package insertion_sort

import (
	"Algorithm/test/data"
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testData := data.RandomData(1000)
	fmt.Println(InsertionSort(testData))
}
