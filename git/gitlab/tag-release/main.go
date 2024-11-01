package main

import (
	"fmt"
	"os"

	"github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v3"
)

type Tag struct {
	Project string `yaml:"project"`
	Branch  string `yaml:"branch"`
	TagName string `yaml:"tagName"`
}

type Config struct {
	Tags []Tag `yaml:"tags"`
}

func main() {
	server := ""
	pat := ""

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

	git, err := gitlab.NewClient(pat, gitlab.WithBaseURL(server))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	pipelineOps := &gitlab.ListProjectPipelinesOptions{Ref: gitlab.Ptr(config.Tags[0].TagName)}
	pipelines, _, err := git.Pipelines.ListProjectPipelines(config.Tags[0].Project, pipelineOps)
	if err != nil {

		return
	}

	for i, pipeline := range pipelines {
		fmt.Println(i, pipeline.ID, pipeline.Ref, pipeline.Status, pipeline.Source, pipeline.WebURL, pipeline.IID)
	}

	jobs, _, err := git.Jobs.ListPipelineJobs(config.Tags[0].Project, pipelines[0].ID, &gitlab.ListJobsOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, job := range jobs {
		if job.Name == "package" {
			fmt.Println(i, " id -> ", job.ID, " name -> ", job.Name, " status -> ", job.Status, " tag -> ", job.Tag, " url -> ", job.WebURL, " state -> ", job.Stage, "pipeline job status ->", job.Pipeline.Status)
		}

	}

	//for _, tag := range config.Tags {
	//	tagCreatOpts := &gitlab.CreateTagOptions{
	//		TagName: gitlab.Ptr(tag.TagName),
	//		Ref:     gitlab.Ptr(tag.Branch),
	//		Message: gitlab.Ptr("Tag created from branch " + tag.Branch),
	//	}
	//
	//	t, _, err := git.Tags.CreateTag(tag.Project, tagCreatOpts)
	//	if err != nil {
	//		fmt.Printf("Failed to create tag: %v\n", err)
	//		os.Exit(1)
	//	}
	//
	//	fmt.Printf("Tag created successfully: %s\n", t.Name)
	//}

}
