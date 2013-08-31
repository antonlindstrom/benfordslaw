package main

import (
	"fmt"
	"math"
)

// Return the leading digit of an integer
func leadingDigit(num int) int {
	fnum := float64(num)

	zeroes := math.Floor(math.Log10(fnum))
	leading := fnum * math.Pow10(-int(zeroes))

	return int(leading)
}

// Calculate probability for digit d
func benfordProbability(d int) float64 {
	return (math.Log10(1.0+(1.0/float64(d))) * 100)
}

// Run through calculations of dataset
func countDataset(numbers []int) (int, []int) {
	count := make([]int, 10)

	for i := range numbers {
		count[leadingDigit(numbers[i])]++
	}

	return int(len(numbers)), count
}

// Calculate percentage
func percentage(count, total int) float64 {
	return float64(count) / float64(total) * 100
}

func main() {
	dataset := []int{1000, 230, 540, 34, 100, 140}

	total, count := countDataset(dataset)

	for i := 1; i < 10; i++ {
		e := benfordProbability(i)
		p := percentage(count[i], total)
		d := p - e

		fmt.Printf("{num: %d, estimate: %f, dataset: %f, diff: %f}\n", i, e, p, d)
	}
}
