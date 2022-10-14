package cmd

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/forbole/bdtool/utils"
)

func Clone(url, directory string, cloneBranch string) *git.Repository {
	// Clone the repository to the given directory
	Info("cloning %s to %s", url, directory)

	repo, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(cloneBranch),
		SingleBranch:  true,
	})
	utils.CheckError(err)

	return repo
}
