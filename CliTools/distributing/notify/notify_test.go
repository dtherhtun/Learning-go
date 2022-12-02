//go:build !integration

package notify

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		s Severity
	}{
		{SeverityLow},
		{SeverityNormal},
		{SeverityUrgent},
	}

	for _, tc := range testCases {
		name := tc.s.String()
		expMessage := "Message"
		expTitle := "Title"
		t.Run(name, func(t *testing.T) {
			n := New(expTitle, expMessage, tc.s)

			if n.message != expMessage {
				t.Errorf("Expected %q, got %q instead\n", expMessage, n.message)
			}

			if n.title != expTitle {
				t.Errorf("Expected %q, got %q instead\n", expTitle, n.title)
			}

			if n.severity != tc.s {
				t.Errorf("Expected %q, got %q instead\n", tc.s, n.severity)
			}
		})
	}
}

func TestSeverityString(t *testing.T) {
	testCases := []struct {
		s   Severity
		exp string
		os  string
	}{
		{SeverityLow, "low", "linux"},
		{SeverityNormal, "normal", "linux"},
		{SeverityUrgent, "critical", "linux"},
		{SeverityLow, "Low", "darwin"},
		{SeverityNormal, "Normal", "darwin"},
		{SeverityUrgent, "Critical", "darwin"},
		{SeverityLow, "Info", "windows"},
		{SeverityNormal, "Warning", "windows"},
		{SeverityUrgent, "Error", "windows"},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s%d", tc.os, tc.s)
		t.Run(name, func(t *testing.T) {
			if runtime.GOOS != tc.os {
				t.Skip("Skipped: not OS", runtime.GOOS)
			}
			sev := tc.s.String()
			if sev != tc.exp {
				t.Errorf("Expected %q,got %q instead\n", tc.exp, sev)
			}
		})
	}
}
