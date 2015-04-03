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
func (d *Dataset) ReadCSV(filename string) error {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Printf("File open error: %v\n", err)
		return err
	}

	reader := csv.NewReader(file)
	records, errCSV := reader.ReadAll()

	if errCSV != nil {
		log.Printf("CSV Error: %v\n", err)
		return err
	}

	data := make([]int, len(records))

	for i, line := range records {
		converted, e := strconv.Atoi(line[0])

		if e != nil {
			log.Printf("Error converting value from CSV: %v\n", e)
			return err
		}

		data[i] = converted
	}

	d.collection = data
	return nil
}

// Wrapper for ReadCSV()
func LoadCSV(filename string) ([]int, error) {
	dataset := new(Dataset)
	err := dataset.ReadCSV(filename)
	if err != nil {
		return []int{}, err
	}
	return dataset.collection, nil
}
