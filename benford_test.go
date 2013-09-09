package main

import (
	"log"
	"testing"
)

func TestJsonString(t *testing.T) {
	var BenfordDigit = BenfordDigit{1, 14, 30.103, 25.454545}

	resultString := "{\"LeadingDigit\":1,\"Count\":14,\"Estimate\":30.103,\"Dataset\":25.454545}"

	if BenfordDigit.JsonString() != resultString {
		t.Error("Unexpected JSON String!")
		log.Printf("Result should have been %s was %s\n", resultString, BenfordDigit.JsonString())
	}
}
