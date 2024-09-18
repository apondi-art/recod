package debit

import "testing"

func TestDebit(t *testing.T) {
	input := 1000
	val := 100
	expected := 1100
	output := Debit(input, val)
	if expected != output {
		t.Errorf("wanted %d but got %d", expected, output)
	}
}
