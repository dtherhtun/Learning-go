//go:build integration

package notify

import (
	"testing"

	"github.com/dtherhtun/Learning-go/CliTools/distributing/notify"
)

func TestSend(t *testing.T) {
	n := notify.New("test title", "test msg", SeverityNormal)

	err := n.Send()

	if err != nil {
		t.Error(err)
	}
}
