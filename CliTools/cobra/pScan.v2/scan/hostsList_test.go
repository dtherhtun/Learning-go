package scan_test

import (
	"errors"
	"github.com/dtherhtun/Learning-go/CliTools/cobra/pScan.v2/scan"
	"os"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"AddNew", "host2", 2, nil},
		{"AddExisting", "host1", 1, scan.ErrExists},
	}

	for testId, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("\tStatus\tTest %s: When Adding %s to hosts list.", tc.name, tc.host)
			hl := &scan.HostsList{}

			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}

			err := hl.Add(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("\t%s\tTest %d:\tExpected error, got nil instead\n", failed, testId)
				}
				t.Logf("\t%s\tTest %d:\tNot nil as expected", succeed, testId)

				if !errors.Is(err, tc.expectErr) {
					t.Fatalf("\t%s\tTest %d:\tExpected error %q, got %q instead\n", failed, testId, tc.expectErr, err)
				}
				t.Logf("\t%s\tTest %d:\tGot %q error as expected", succeed, testId, tc.expectErr)
				return
			}
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tExpected no error, got %q instead\n", failed, testId, err)
			}
			t.Logf("\t%s\tTest %d:\tNo error as expected", succeed, testId)

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("\t%s\tTest %d:\tExpected list length %d, got %d instead\n", failed, testId, tc.expectLen, len(hl.Hosts))
			}
			t.Logf("\t%s\tTest %d:\tGot correct hosts list as expected", succeed, testId)

			if hl.Hosts[1] != tc.host {
				t.Errorf("\t%s\tTest %d:\tExpected host name %q as index 1, got %q instead\n", failed, testId, tc.host, hl.Hosts[1])
			}
			t.Logf("\t%s\tTest %d:\tGot correct hosts name as expected", succeed, testId)
		})
	}
}

func TestSaveLoad(t *testing.T) {
	t.Logf("\tStatus\tTest\tWhen checking status of saved and loaded hosts list file")
	hl1 := scan.HostsList{}
	hl2 := scan.HostsList{}

	hostname := "host1"
	hl1.Add(hostname)

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("\t%s\tTest:\tError saving list to file: %s", failed, err)
	}
	t.Logf("\t%s\tTest:\tSuccessfully saved hosts list file as expected.", succeed)

	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("\t%s\tTest:\tError getting list from file: %s", failed, err)
	}
	t.Logf("\t%s\tTest:\tSuccessfully Loaded hosts list file as expected.", succeed)

	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Errorf("\t%s\tTest:\tHost %q should match %q host.", failed, hl1.Hosts[0], hl2.Hosts[0])
	}
	t.Logf("\t%s\tTest:\tSaved file and Loaded file are equeal as expected.", succeed)
}

func TestLoadNoFile(t *testing.T) {
	t.Logf("\tStatus\tTest\tWhen checking status of loaded no file")

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	if err := os.Remove(tf.Name()); err != nil {
		t.Fatalf("Error deleting temp file: %s", err)
	}

	hl := scan.HostsList{}
	if err := hl.Load(tf.Name()); err != nil {
		t.Errorf("\t%s\tTest:\tExpected no error, got %q instead\n", failed, err)
	}
	t.Logf("\t%s\tTest:\tNo error as Expected", succeed)
}
