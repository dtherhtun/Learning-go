package pork

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dtherhtun/Learning-go/ops/buildingOpsTools/nap"
	"github.com/spf13/cobra"
)

type ForkResponse struct {
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a GitHub Repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply a repository")
		}
		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork repository: ", err)
		}
	},
}

func ForkRepository(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("repository must be in the format owner/project")
	}
	return GitHubAPI().Call("fork", map[string]string{
		"owner": values[0],
		"repo":  values[1],
	})
}

func ForkSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Forked to repository: %s\n", respContent.FullName)
	return nil
}

func GetForkResource() *nap.RestResource {
	forkRouter := nap.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess)
	forkRouter.RegisterFunc(401, func(_ *http.Response, _ interface{}) error {
		return fmt.Errorf("you must set an authentication token")
	})
	fork := nap.NewResource("/repos/{{.owner}}/{{.repo}}/forks", "POST", forkRouter)
	return fork
}
