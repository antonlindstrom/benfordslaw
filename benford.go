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
	c := counter.Process(loader.LoadCSV())

	for i := range c.Digits {
		var m BenfordDigit = BenfordDigit{
			c.Digits[i].Leading,
			c.Digits[i].Count,
			counter.BenfordProbability(i),
			counter.Percentage(i, c.Total),
		}

		b, err := json.Marshal(m)

		if err != nil {
			fmt.Printf("Failed to Marshal JSON: %s\n", err)
		}

		fmt.Printf("%s\n", string(b[:]))
	}
}
