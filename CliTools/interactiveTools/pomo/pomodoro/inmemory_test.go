package pomodoro_test

import (
	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro"
	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro/repository"
	"testing"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
