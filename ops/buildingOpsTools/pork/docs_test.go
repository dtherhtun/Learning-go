package pork

import "testing"

func TestGetRepositoryReadme(t *testing.T) {
	content := GetRepositoryReadme("myrepository")
	if content != nil {
		t.Fail()
	}
}
