package loader

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Dataset struct {
	collection []int
}

// Load dataset from CSV file
func (d *Dataset) ReadCSV(filename string) {
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

// Wrapper for ReadCSV()
func LoadCSV(filename string) ([]int) {
	dataset := new(Dataset)
	dataset.ReadCSV(filename)
	return dataset.collection
}
