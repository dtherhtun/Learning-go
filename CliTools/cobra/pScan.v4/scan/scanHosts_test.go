package scan_test

import (
	"github.com/dtherhtun/Learning-go/CliTools/cobra/pScan/scan"
	"net"
	"strconv"
	"testing"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}

	if ps.Open.String() != "closed" {
		t.Errorf("\t%s\tTest:\tExpected %q, got %q instead\n", failed, "closed", ps.Open.String())
	} else {
		t.Logf("\t%s\tTest:\tPort is closed as expect.", succeed)
	}

	ps.Open = true

	if ps.Open.String() != "open" {
		t.Errorf("\t%s\tTest:\tExpected %q, got %q instead\n", failed, "open", ps.Open.String())
	} else {
		t.Logf("\t%s\tTest:\tPort is open as expect.", succeed)
	}
}

func TestRunHostFound(t *testing.T) {
	testCases := []struct {
		name        string
		expectState string
	}{
		{"OpenPort", "open"},
		{"ClosedPort", "closed"},
	}

	host := "localhost"
	hl := &scan.HostsList{}
	hl.Add(host)

	ports := []int{}

	for _, tc := range testCases {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, "0"))
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal(err)
		}

		ports = append(ports, port)

		if tc.name == "ClosedPort" {
			ln.Close()
		}
	}

	res := scan.Run(hl, ports)

	if len(res) != 1 {
		t.Fatalf("\t%s\tTEST:\tExpected 1 results, go %d instead\n", failed, len(res))
	}

	if res[0].Host != host {
		t.Fatalf("\t%s\tTEST:\tExpected host %q, got %q instead\n", failed, host, res[0].Host)
	}

	if res[0].NotFound {
		t.Fatalf("\t%s\tTEST:\tExpected host %q to be found\n", failed, host)
	}

	if len(res[0].PortStates) != 2 {
		t.Fatalf("\t%s\tTEST:\tExpected 2 port states, got %d instead\n", failed, len(res[0].PortStates))
	}

	for i, tc := range testCases {
		if res[0].PortStates[i].Port != ports[i] {
			t.Errorf("Expected port %d, got %d instead\n", ports[i], res[0].PortStates[i].Port)
		}
		if res[0].PortStates[i].Open.String() != tc.expectState {
			t.Errorf("Expected port %d to be %s\n", ports[i], tc.expectState)
		}
	}
}

func TestRunHostNotFound(t *testing.T) {
	host := "389.389.389.389"
	hl := &scan.HostsList{}
	hl.Add(host)

	res := scan.Run(hl, []int{})

	if len(res) != 1 {
		t.Fatalf("Expected 1 results, got %d instead\n", len(res))
	}

	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q instead\n", host, res[0].Host)
	}

	if !res[0].NotFound {
		t.Errorf("Expected host %q NOT to be found\n", host)
	}

	if len(res[0].PortStates) != 0 {
		t.Fatalf("Expected 0 port states, got %d instead\n", len(res[0].PortStates))
	}
}
