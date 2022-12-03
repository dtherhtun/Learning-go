//go:build inmemory && containers

package cmd

import (
	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro"
	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
