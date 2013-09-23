package main

import (
	"github.com/antonlindstrom/benfordslaw"
	"github.com/antonlindstrom/benfordslaw/loader"
	"log"
)

const FILENAME string = "./data/FlowOfFunds.csv"

func main() {
	csv := loader.LoadCSV(FILENAME)
	log.Printf("%s\n", benfordslaw.Process(csv))
}
