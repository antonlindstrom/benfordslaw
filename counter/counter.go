package counter

import (
	"math"
)

// Return the leading digit of an integer
func LeadingDigit(num int) int {
	fnum := float64(num)

	zeroes := math.Floor(math.Log10(fnum))
	leading := fnum * math.Pow10(-int(zeroes))

	return int(leading)
}

// Calculate probability for digit d
func BenfordProbability(d int) float64 {
	return (math.Log10(1.0+(1.0/float64(d))) * 100)
}

// Calculate percentage
func Percentage(count, total int) float64 {
	return float64(count) / float64(total) * 100
}

// Run through calculations of dataset
func Process(numbers []int) (int, []int) {
	count := make([]int, 10)

	for i := range numbers {
		leadingDigit := LeadingDigit(numbers[i])
		count[leadingDigit]++
	}

	return len(numbers), count
}
