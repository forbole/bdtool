package cmd

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/forbole/bdtool/utils"
)

func Push(repo *git.Repository, GitHubToken string) {
	Info("Pushing to remote repo")

	err := repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "-",
			Password: GitHubToken,
		},
		InsecureSkipTLS: true,
	})

	if err != nil {
		utils.CheckError(err)
	}

}
