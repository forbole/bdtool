package main

import (
	gitrepo "github.com/forbole/bdtool/types/repo"
	"github.com/forbole/bdtool/utils"
)

func main() {
	// Ask for neccessary info
	chainInfo := utils.GetChainInfo()
	gitConfig := utils.GetGitConfig()
	chainConfig := utils.GetChainConfig(chainInfo)

	repo := gitrepo.New(chainInfo, chainConfig, gitConfig)

	// Execute "git checkout -b chains/{chain name}/{chain type}"
	repo.Checkout()

	// Write config file & Copy chain's icon and logo
	repo.WriteConfig().CopyImages()

	// Commit, push, and open PR
	repo.Commit().Push().PullRequest()

	// Remove temporal repo when all is done
	repo.Remove()
}
