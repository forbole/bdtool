package cmd

import (
	"fmt"
	"time"

	"github.com/forbole/bdtool/types"
	"github.com/forbole/bdtool/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func Commit(chain *types.ChainInfo, repo *git.Repository, iconFileName, logoFileName string) {
	w, err := repo.Worktree()
	utils.CheckError(err)

	Info("git add & commit: %s-%s.json, %s, and %s", chain.Name, chain.Type, iconFileName, logoFileName)

	// git add config file
	cfgFile := fmt.Sprintf("src/configs/chain_configs/%s-%s.json", chain.Name, chain.Type)
	_, err = w.Add(cfgFile)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while adding %s: %s", cfgFile, err))
	}

	// git add icon file
	iconFile := fmt.Sprintf("public/images/%s/%s", chain.Name, iconFileName)
	_, err = w.Add(iconFile)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while adding %s: %s", iconFile, err))
	}

	// git add logo file
	logoFile := fmt.Sprintf("public/images/%s/%s", chain.Name, logoFileName)
	_, err = w.Add(logoFile)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while adding %s: %s", logoFile, err))
	}

	Info("performing git commit, please enter Author Name and Email:")

	author, err := utils.GetGitConfig("user.name")
	if err != nil {
		author = utils.GetInput("Author Name")
	}

	email, err := utils.GetGitConfig("user.email")
	if err != nil {
		email = utils.GetInput("Author Email")
	}

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
