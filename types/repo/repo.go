package repo

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-git/go-git/v5"
	gitcmd "github.com/forbole/bdtool/gitcmd"
	"github.com/forbole/bdtool/types"
	"github.com/forbole/bdtool/utils"
)

type Repo struct {
	Repo           *git.Repository
	ChainInfo      *types.Chain
	ChainConfig    *types.ChainConfig
	Path           string
	GitHubToken    string
	PrTargetBranch string
}

func New(repoURL, cloneBranch, prTargetBranch string, chainInfo *types.Chain, chainConfig *types.ChainConfig, GitHubToken string) *Repo {
	// Prepare file destination
	path, err := prepareFileDest()
	if err != nil {
		utils.CheckError(err)
	}

	// Clone the repo
	repo := gitcmd.Clone(repoURL, path, cloneBranch)

	return &Repo{
		Repo:           repo,
		ChainConfig:    chainConfig,
		ChainInfo:      chainInfo,
		Path:           path,
		GitHubToken:    GitHubToken,
		PrTargetBranch: prTargetBranch,
	}
}

func (r *Repo) Checkout() *Repo {
	newBranch := fmt.Sprintf("chains/%s/%s", r.ChainInfo.Name, r.ChainInfo.Type)
	branchRef := gitcmd.Branch(r.Repo, newBranch)
	gitcmd.Checkout(r.Repo, branchRef.Name())
	return r
}

func (r *Repo) WriteConfig() *Repo {
	file := fmt.Sprintf("%s/src/configs/chain_configs/%s-%s.json", r.Path, r.ChainInfo.Name, r.ChainInfo.Type)

	_, err := os.Create(file)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while creating chain config file: %s", err))
	}

	err = ioutil.WriteFile(file, utils.GetConfigBz(r.ChainConfig), 0600)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while writing chain config: %s", err))
	}

	return r
}

func (r *Repo) CopyImages() *Repo {
	imgDir := fmt.Sprintf("%s/public/images/%s", r.Path, r.ChainInfo.Name)
	err := os.MkdirAll(imgDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Get icon name and path
	iconPath := utils.GetInput("Enter icon file path")
	iconFileName := getFileNameFromPath(iconPath)

	// Copy icon
	copy(iconPath, fmt.Sprintf("%s/%s", imgDir, iconFileName))

	// Get logo name and path
	logoPath := utils.GetInput("Enter logo file path")
	logoFileName := getFileNameFromPath(logoPath)

	// Copy logo
	copy(logoPath, fmt.Sprintf("%s/%s", imgDir, logoFileName))

	return r
}

func (r *Repo) Commit() *Repo {
	gitcmd.Commit(r.ChainInfo, r.Repo)
	return r
}

func (r *Repo) Push() *Repo {
	gitcmd.Push(r.Repo, r.GitHubToken)
	return r
}

func (r *Repo) PullRequest() *Repo {
	err := gitcmd.PullRequest(r.ChainInfo, r.PrTargetBranch, r.GitHubToken)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while opening pull request : %s", err))
	}

	return r
}

func (r *Repo) RemoveDir() {
	remove := utils.GetBool("remove temporal directory")
	if !remove {
		return
	}

	err := os.RemoveAll(r.Path)
	if err != nil {
		utils.CheckError(fmt.Errorf("error while removing temp_BD directory : %s", err))
	}
}
