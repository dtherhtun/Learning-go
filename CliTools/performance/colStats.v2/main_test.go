package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name   string
		col    int
		op     string
		exp    string
		files  []string
		expErr error
	}{
		{
			name:   "RunAvg1File",
			col:    3,
			op:     "avg",
			exp:    "227.6\n",
			files:  []string{"./testdata/example.csv"},
			expErr: nil,
		},
		{
			name:   "RunAvgMultiFiles",
			col:    3,
			op:     "avg",
			exp:    "233.84\n",
			files:  []string{"./testdata/example.csv", "./testdata/example2.csv"},
			expErr: nil,
		},
		{
			name:   "RunFailRead",
			col:    2,
			op:     "avg",
			exp:    "",
			files:  []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			expErr: os.ErrNotExist,
		},
		{
			name:   "RunFailColumn",
			col:    0,
			op:     "avg",
			exp:    "",
			files:  []string{"./testdata/example.csv"},
			expErr: ErrInvaildColumn,
		},
		{
			name:   "RunFailNoFiles",
			col:    2,
			op:     "avg",
			exp:    "",
			files:  []string{},
			expErr: ErrNoFiles,
		},
		{
			name:   "RunFailOperation",
			col:    2,
			op:     "invalid",
			exp:    "",
			files:  []string{"./testdata/example.csv"},
			expErr: ErrInvalidOperation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &res)

			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error. Got nil instead")
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error %q, got %q instead", tc.expErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}
			if res.String() != tc.exp {
				t.Errorf("Expected %q, got %q instead", tc.exp, &res)
			}
		})
	}
}

func BenchmarkRun(b *testing.B) {
	filenames, err := filepath.Glob("./testdata/benchmark/*.csv")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := run(filenames, "avg", 2, ioutil.Discard); err != nil {
			b.Error(err)
		}
	}
}
