package main

import (
	gitrepo "github.com/forbole/bdtool/types/repo"
	"github.com/forbole/bdtool/utils"
)

var (
	REPO_URL         = "https://github.com/forbole/big-dipper-2.0-cosmos"
	CLONE_BRANCH     = "refs/heads/bdu-585-improve-setup-process"
	PR_TARGET_BRANCH = "refs/heads/bdu-585-improve-setup-process-clone"
)

func main() {
	// Ask for neccessary info
	chainInfo := utils.GetChainInfo()
	GitHubToken := utils.GetTokenInput("GitHub Personal Access Token")
	chainConfig := utils.GetConfig(chainInfo)

	repo := gitrepo.New(
		REPO_URL, CLONE_BRANCH, PR_TARGET_BRANCH,
		chainInfo, chainConfig, GitHubToken,
	)

	// Execute "git checkout -b chains/{chain name}/{chain type}"
	repo.Checkout()

	// Write config file; Copy chain's icon and logo
	repo.WriteConfig().CopyImages()

	// Commit, push, and open PR
	repo.Commit().Push().PullRequest()

	// Remove temporal repo when all is done
	repo.RemoveDir()
}
