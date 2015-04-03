package main

import (
	"github.com/antonlindstrom/benfordslaw"
	"github.com/antonlindstrom/benfordslaw/loader"
	"log"
)

const FILENAME string = "./data/FlowOfFunds.csv"

func main() {
	csv, err := loader.LoadCSV(FILENAME)
	if err != nil {
		log.Printf("Failed to load CSV: %s\n", err)
		return
	}
	log.Printf("%s\n", benfordslaw.Process(csv))
}
