package main

import (
	"github.com/antonlindstrom/benford/loader"
	"log"
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

// Run through calculations of dataset
func CountDataset(numbers []int) (int, []int) {
	count := make([]int, 10)

	for i := range numbers {
		count[LeadingDigit(numbers[i])]++
	}

	return int(len(numbers)), count
}

// Calculate percentage
func Percentage(count, total int) float64 {
	return float64(count) / float64(total) * 100
}

func main() {
	total, count := CountDataset(loader.LoadCSV())

	for i := 1; i < 10; i++ {
		e := BenfordProbability(i)
		p := Percentage(count[i], total)
		d := p - e

		log.Printf("{num: %d, count: %d estimate: %f%%, dataset: %f%%, diff: %f%%}\n", i, count[i], e, p, d)
	}
	log.Printf("{total_count: %d}\n", total)
}
