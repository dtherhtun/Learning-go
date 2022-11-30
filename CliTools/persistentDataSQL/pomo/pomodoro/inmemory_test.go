//go:build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro"
	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
