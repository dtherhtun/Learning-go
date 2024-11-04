package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v3"
)

type Tag struct {
	Project string `yaml:"project"`
	Branch  string `yaml:"branch"`
	TagName string `yaml:"tagName"`
}

type Config struct {
	Tags      []Tag  `yaml:"tags"`
	Server    string `yaml:"server"`
	PAT       string `yaml:"pat"`
	TargetEnv string `yaml:"targetEnv"`
}

func main() {

	data, err := os.ReadFile("tags.yaml")
	if err != nil {
		fmt.Printf("Failed to read YAML file: %v\n", err)
		return
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("Failed to parse YAML file: %v\n", err)
		return
	}

	git, err := gitlab.NewClient(config.PAT, gitlab.WithBaseURL(config.Server))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for _, tag := range config.Tags {
		tagCreatOpts := &gitlab.CreateTagOptions{
			TagName: gitlab.Ptr(tag.TagName),
			Ref:     gitlab.Ptr(tag.Branch),
			Message: gitlab.Ptr("Tag created from branch " + tag.Branch),
		}

		t, _, err := git.Tags.CreateTag(tag.Project, tagCreatOpts)
		if err != nil {
			fmt.Printf("Failed to create tag: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Tag created successfully: %s\n", t.Name)
		time.Sleep(2 * time.Second)
	}

	// I would like to use concurrency start from here
	for _, tag := range config.Tags {
		pipelineOps := &gitlab.ListProjectPipelinesOptions{Ref: gitlab.Ptr(tag.TagName)}
		pipelines, _, err := git.Pipelines.ListProjectPipelines(tag.Project, pipelineOps)
		if err != nil {
			return
		}

		for {
			pkgFinished := false
			jobs, _, err := git.Jobs.ListPipelineJobs(tag.Project, pipelines[0].ID, &gitlab.ListJobsOptions{})
			if err != nil {
				fmt.Println(err)
				return
			}

			var playJobID int

			for _, job := range jobs {
				if job.Name == config.TargetEnv {
					playJobID = job.ID
				}
				if job.Name == "package" && job.Status == "success" {
					pkgFinished = true
					break
				}
			}

			if pkgFinished {
				fmt.Println("Package finished successfully")
				job, _, err := git.Jobs.PlayJob(tag.Project, playJobID, &gitlab.PlayJobOptions{})
				if err != nil {
					return
				}
				fmt.Println(job.Name, " is running.")
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
}
