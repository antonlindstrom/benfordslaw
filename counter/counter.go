package counter

import (
	"math"
)

type Digit struct {
	Leading   int
	Count     int
}

type Collection struct {
	Digits  map[int]Digit
	Total   int
}

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
func (c *Collection) CountDataset(numbers []int) {
	c.Digits = make(map[int]Digit)
	c.Total = len(numbers)

	for i := range numbers {
		leading := LeadingDigit(numbers[i])

		var currentDigit Digit
		count := c.Digits[leading].Count+1

		currentDigit = Digit{
			Leading:leading,
			Count:count,
		}

		c.Digits[leading] = currentDigit
	}
}

func Process(numbers []int) (*Collection) {
	c := new(Collection)
	c.CountDataset(numbers)
	return c
}
