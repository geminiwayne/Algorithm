package data

import (
	"math/rand"
	"time"
)

// create int array for test
func RandomData(length int) []int {
	var testData []int
	rand.Seed(time.Now().Unix())
	count := 0
	for count < length {
		testData = append(testData, rand.Intn(length))
		count++
	}
	return testData
}
