package pork

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	destRepo           string
	sourceRepo         string
	pullRequestTitle   string
	pullRequestMessage string
)

type PullRequestPayload struct {
	Title        string `json:"title"`
	Message      string `json:"body"`
	SourceBranch string `json:"head"`
	DestBranch   string `json:"base"`
	Modify       bool   `json:"maintainer_can_modify"`
}

type PullRequestResponse struct {
	URL string `json:"html_url"`
}

var PullRequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Short: "Create a Pull Request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreatePullRequest(); err != nil {
			log.Fatalln("Failed to create pull request:", err)
		}
	},
}

func CreatePullRequest() error {
	sourceValues := strings.Split(sourceRepo, ":")
	if !(len(sourceValues) == 1 || len(sourceValues) == 2) {
		return fmt.Errorf("source repository must in the format [owner:]branch got %v", sourceRepo)
	}
	destBranchValues := strings.Split(destRepo, ":")
	if len(destBranchValues) != 2 {
		return fmt.Errorf("destination repository must be in the format owner/project:branch got %v", destRepo)
	}
	destValues := strings.Split(destBranchValues[0], "/")
	if len(destValues) != 2 {
		return fmt.Errorf("destination repository must be in the format owner/project:branch got %v", destRepo)
	}
	payload := &PullRequestPayload{
		Title:        pullRequestTitle,
		Message:      pullRequestMessage,
		SourceBranch: sourceRepo,
		DestBranch:   destBranchValues[1],
		Modify:       true,
	}
	return GitHubAPI().Call("pullrequest", map[string]string{
		"owner":   destValues[0],
		"project": destValues[1],
	})
}
