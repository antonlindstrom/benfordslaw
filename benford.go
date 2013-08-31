package main

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

type Dataset struct {
	collection []int
}

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

// Load dataset from CSV file
func (d *Dataset) loadData(filename string) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Printf("File open error: %v\n", err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	records, errCSV := reader.ReadAll()

	if errCSV != nil {
		log.Printf("CSV Error: %v\n", err)
		os.Exit(1)
	}

	data := make([]int, len(records))

	for i, line := range records {
		converted, e  := strconv.Atoi(line[0])

		if e != nil {
			log.Printf("Error converting value from CSV: %v\n", e)
			os.Exit(2)
		}

		data[i] = converted
	}

	d.collection = data
}

func main() {
	dataset := new(Dataset)
	dataset.loadData("./data/FlowOfFunds.csv")

	total, count := countDataset(dataset.collection)

	for i := 1; i < 10; i++ {
		e := benfordProbability(i)
		p := percentage(count[i], total)
		d := p - e

		log.Printf("{num: %d, estimate: %f%%, dataset: %f%%, diff: %f%%}\n", i, e, p, d)
	}
}
