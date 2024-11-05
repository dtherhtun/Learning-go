package main

import (
	"fmt"
	"os"
	"sync"
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
		fmt.Printf("[-] Failed to read YAML file: %v\n", err)
		return
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("[-] Failed to parse YAML file: %v\n", err)
		return
	}

	git, err := gitlab.NewClient(config.PAT, gitlab.WithBaseURL(config.Server))
	if err != nil {
		fmt.Printf("[-] Failed to create client: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	tagCh := make(chan Tag)

	go func() {
		for tag := range tagCh {
			tagCreatOpts := &gitlab.CreateTagOptions{
				TagName: gitlab.Ptr(tag.TagName),
				Ref:     gitlab.Ptr(tag.Branch),
				Message: gitlab.Ptr("Tag created from branch " + tag.Branch),
			}

			t, _, err := git.Tags.CreateTag(tag.Project, tagCreatOpts)
			if err != nil {
				fmt.Printf("[-] Failed to create tag: %v\n", err)
				return
			}

			fmt.Printf("[+] Tag created successfully: %s\n", t.Name)
			time.Sleep(2 * time.Second)
			wg.Add(1)
			go searchAndTriggerJobFromRelativePipeline(git, tag, t.Name, config.TargetEnv, &wg)
		}
	}()

	for _, tag := range config.Tags {
		tagCh <- tag
	}
	close(tagCh)

	wg.Wait()
}

func searchAndTriggerJobFromRelativePipeline(git *gitlab.Client, tag Tag, tagName string, targetEnv string, wg *sync.WaitGroup) {
	defer wg.Done()

	pipelineOps := &gitlab.ListProjectPipelinesOptions{Ref: gitlab.Ptr(tagName)}
	pipelines, _, err := git.Pipelines.ListProjectPipelines(tag.Project, pipelineOps)
	if err != nil {
		fmt.Printf("[-] Failed to list pipelines for tag %s: %v\n", tagName, err)
		return
	}

	if len(pipelines) == 0 {
		fmt.Printf("[-] No pipelines found for tag %s\n", tagName)
		return
	}

	for {
		pkgFinished := false
		jobs, _, err := git.Jobs.ListPipelineJobs(tag.Project, pipelines[0].ID, &gitlab.ListJobsOptions{})
		if err != nil {
			fmt.Printf("[-] Failed to list jobs for tag %s: %v\n", tagName, err)
			return
		}

		var playJobID int

		for _, job := range jobs {
			if job.Name == targetEnv {
				playJobID = job.ID
			}
			if job.Name == "package" && job.Status == "success" {
				pkgFinished = true
				fmt.Printf("[+] Pipeline job %s finished done\n", job.Name)
				break
			}
		}

		if pkgFinished {
			job, _, err := git.Jobs.PlayJob(tag.Project, playJobID, &gitlab.PlayJobOptions{})
			if err != nil {
				fmt.Printf("Failed to play job for tag %s: %v\n", tagName, err)
				return
			}
			fmt.Printf("[+] %s - %s job is running", job.Project.Name, job.Name)
			break
		}
		time.Sleep(10 * time.Second)
	}
}
