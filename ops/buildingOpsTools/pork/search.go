package pork

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for GitHub repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		repoList := SearchByKeyword(args)
		for _, repo := range repoList {
			fmt.Println(repo)
		}
	},
}

func SearchByKeyword(keywords []string) []string {
	return []string{"myrepo"}
}
