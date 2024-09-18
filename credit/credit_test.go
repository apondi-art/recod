package credit

import "testing"

func TestCredit(t *testing.T) {
	input := 1000
	val := 100
	expected := 900
	output := Credit(input, val)
	if expected != output {
		t.Errorf("wanted %d but got %d", expected, output)
	}
}
