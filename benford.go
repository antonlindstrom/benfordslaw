package main

import (
	"encoding/json"
	"github.com/antonlindstrom/benford/counter"
	"github.com/antonlindstrom/benford/loader"
	"fmt"
)

type BenfordDigit struct {
	LeadingDigit   int
	Count          int
	Estimate       float64
	Dataset        float64
}

func main() {
	numbers      := loader.LoadCSV()
	total, count := counter.Process(numbers)

	for i := 1; i < 10; i++ {
		var m BenfordDigit = BenfordDigit{
			i,
			count[i],
			counter.BenfordProbability(i),
			counter.Percentage(i, total),
		}

		b, err := json.Marshal(m)

		if err != nil {
			fmt.Printf("Failed to Marshal JSON: %s\n", err)
		}

		fmt.Printf("%s\n", string(b[:]))
	}
}
