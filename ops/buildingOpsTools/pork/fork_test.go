package pork

import "testing"

func TestForkRepository(t *testing.T) {
	if err := ForkRepository("myrepo"); err != nil {
		t.Fail()
	}
}
