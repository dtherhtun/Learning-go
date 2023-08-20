package pork

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type GitHubRepo struct {
	RepoDir string
	owner   string
	project string
	repo    *git.Repository
}

func NewGitHubRepo(repository string) (*GitHubRepo, error) {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return nil, fmt.Errorf("repository name must be in format owner/project")
	}
	return &GitHubRepo{
		owner:   values[0],
		project: values[1],
	}, nil
}

func (g *GitHubRepo) RepositoryURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.owner, g.project)
}

func (g *GitHubRepo) Clone(dest string) error {
	fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
	repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL: g.RepositoryURL(),
	})
	if err != nil {
		return err
	}
	g.repo = repo
	g.RepoDir = fullPath

	return nil
}

func (g *GitHubRepo) Checkout(ref string, create bool) error {
	opts := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", ref)),
		Create: create,
	}
	if create {
		head, err := g.repo.Head()
		if err != nil {
			return err
		}
		opts.Hash = head.Hash()
	}
	tree, err := g.repo.Worktree()
	if err != nil {
		return err
	}
	if err := tree.Checkout(opts); err != nil {
		return err
	}
	return nil
}

func (g *GitHubRepo) AddUpStream(repository *GitHubRepo) error {
	_, err := g.repo.CreateRemote(&config.RemoteConfig{
		URLs: []string{repository.RepositoryURL()},
	})
	return err
}
