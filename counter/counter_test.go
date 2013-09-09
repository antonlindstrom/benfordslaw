package counter

import "testing"

func TestLeadingDigit(t *testing.T) {
	if (LeadingDigit(3000) != 3) {
		t.Error("expected 3 from 3000")
	}

	if (LeadingDigit(15) != 1) {
		t.Error("expected 1 from 15")
	}

	if (LeadingDigit(9980) != 9) {
		t.Error("expected 9 of 9980")
	}
}

func TestBenfordProbability(t *testing.T) {
	if (int(BenfordProbability(1)) != 30) {
		t.Error("expected 30 of int 1")
	}

	if (int(BenfordProbability(3)) != 12) {
		t.Error("expected 12 of int 3")
	}
}

func TestProcess(t *testing.T) {
	set := []int{1,1,2}

	total, count := Process(set)

	if (count[1] != 2) {
		t.Error("expected count 2 of 1")
	}

	if (count[2] != 1) {
		t.Error("expected count 1 of 2")
	}

	if (total != 3) {
		t.Error("expected total 3")
	}
}

func TestPercentage(t *testing.T) {
	if (int(Percentage(2,10)) != 20) {
		t.Error("expected 20 of 2/10")
	}
}
