package cmd

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/forbole/bdtool/utils"
)

func Branch(repo *git.Repository, newBranchName string) *plumbing.Reference {
	// Create a new branch to the current HEAD
	Info("creating branch %s", newBranchName)

	headRef, err := repo.Head()
	utils.CheckError(err)

	ref := plumbing.NewHashReference(
		plumbing.NewBranchReferenceName(newBranchName),
		headRef.Hash(),
	)

	// The created reference is saved in the storage.
	err = repo.Storer.SetReference(ref)
	utils.CheckError(err)

	return ref
}
