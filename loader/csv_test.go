package loader

import (
	"testing"
)

func TestLoadCSV(t *testing.T) {
	csv := LoadCSV("../data/FlowOfFunds.csv")

	if csv[0] != 19 {
		t.Error("expected 19 from LoadCSV[0]")
	}

	if csv[1] != 20 {
		t.Error("expected 20 from LoadCSV[1]")
	}

	if csv[54] != 12551 {
		t.Error("expected 12551 from LoadCSV[54]")
	}
}
