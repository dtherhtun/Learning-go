//go:build !inmemory && !containers

package cmd

import (
	"github.com/spf13/viper"

	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro"
	"github.com/dtherhtun/Learning-go/CliTools/interactiveTools/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}

	return repo, nil
}
