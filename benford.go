package benfordslaw

import (
	"encoding/json"
	"github.com/antonlindstrom/benfordslaw/counter"
	"log"
)

type BenfordDigit struct {
	LeadingDigit int
	Count        int
	Estimate     float64
	Dataset      float64
}

type Collection struct {
	Digits []BenfordDigit
}

func (b *Collection) JsonString() string {
	bytes, err := json.Marshal(b)

	if err != nil {
		log.Printf("Failed to Marshal JSON: %s\n", err)
		return ""
	}

	return string(bytes)
}

func (b *Collection) Populate(total int, count []int) {
	b.Digits = make([]BenfordDigit, 10)

	for i := 1; i < 10; i++ {
		var m BenfordDigit = BenfordDigit{
			i,
			count[i],
			counter.BenfordProbability(i),
			counter.Percentage(count[i], total),
		}

		b.Digits[i] = m
	}
}

func Process(set []int) string {
	total, count := counter.Process(set)

	c := new(Collection)
	c.Populate(total, count)

	return c.JsonString()
}
