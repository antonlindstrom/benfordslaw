package main

import (
	"encoding/json"
	"fmt"
	"github.com/antonlindstrom/benford/counter"
	"github.com/antonlindstrom/benford/loader"
)

const FILENAME string = "./data/FlowOfFunds.csv"

type BenfordDigit struct {
	LeadingDigit int
	Count        int
	Estimate     float64
	Dataset      float64
}

func (digit BenfordDigit) JsonString() string {
	b, err := json.Marshal(digit)

	if err != nil {
		fmt.Printf("Failed to Marshal JSON: %s\n", err)
		return ""
	}

	return string(b[:])
}

func main() {
	total, count := counter.Process(loader.LoadCSV(FILENAME))

	for i := 1; i < 10; i++ {
		var m BenfordDigit = BenfordDigit{
			i,
			count[i],
			counter.BenfordProbability(i),
			counter.Percentage(count[i], total),
		}

		fmt.Printf("%s\n", m.JsonString())
	}
}
