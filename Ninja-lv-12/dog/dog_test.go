package dog

import "testing"

type testpair struct {
	hyear int
	dyear int
}

var tests = []testpair{
	{1, 7},
	{2, 14},
	{7, 50},
}

func TestYears(t *testing.T) {
	for _, pair := range tests {
		v := Years(pair.hyear)
		if v != pair.dyear {
			t.Error("For", pair.hyear, "expected", pair.dyear, "got", v)
		}
	}
}
