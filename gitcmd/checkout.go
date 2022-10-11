package cmd

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/forbole/bdtool/utils"
)

func Checkout(r *git.Repository, branch plumbing.ReferenceName) {
	Info("checking out %s", branch)

	w, err := r.Worktree()
	utils.CheckError(err)

	err = w.Checkout(&git.CheckoutOptions{
		Branch: branch,
	})
	utils.CheckError(err)
}
