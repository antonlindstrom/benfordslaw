package benfordslaw

import "testing"

func TestPopulate(t *testing.T) {
	total := 29
	count := []int{0,9,5,4,2,2,2,1,2,1}

	c := new(Collection)
	c.Populate(total, count)

	if (c.Digits[0].Count != 0) {
		t.Fatal("Count of 0 should be 0\n")
	}
}
