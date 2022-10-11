package cmd

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/forbole/bdtool/types"
	"github.com/forbole/bdtool/utils"
)

func Commit(chain *types.Chain, repo *git.Repository) {
	w, err := repo.Worktree()
	utils.CheckError(err)

	Info("git add %s-%s.json", chain.Name, chain.Type)

	filename := fmt.Sprintf("src/configs/chain_configs/%s-%s.json", chain.Name, chain.Type)
	_, err = w.Add(filename)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while adding %s: %s", filename, err))
	}

	Info("committing config file, please enter Author Name and Email:")

	author := utils.GetInput("Author Name")
	email := utils.GetInput("Author Email")

	commitMsg := fmt.Sprintf("add config file for %s %s", chain.Name, chain.Type)
	commit, err := w.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  author,
			Email: email,
			When:  time.Now(),
		},
	})

	utils.CheckError(err)

	// Prints the current HEAD to verify that all worked well.
	obj, err := repo.CommitObject(commit)
	utils.CheckError(err)

	fmt.Println(obj)
}
