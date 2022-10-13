package repo

import (
	"fmt"
	"io/ioutil"
	"os"

	gitcmd "github.com/forbole/bdtool/gitcmd"
	"github.com/forbole/bdtool/types"
	gittypes "github.com/forbole/bdtool/types/git"
	"github.com/forbole/bdtool/utils"
	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Repo         *git.Repository
	ChainInfo    *types.ChainInfo
	ChainConfig  *types.ChainConfig
	Path         string
	GitConfig    *gittypes.GitConfig
	IconFileName string
	LogoFileName string
}

func New(chainInfo *types.ChainInfo, chainConfig *types.ChainConfig, gitConfig *gittypes.GitConfig) *Repo {
	// Prepare file destination
	path, err := prepareFileDest()
	if err != nil {
		utils.CheckError(err)
	}

	// Clone the repo
	repo := gitcmd.Clone(gitConfig.RepoURL, path, gitConfig.CloneBranch)

	return &Repo{
		Repo:        repo,
		ChainConfig: chainConfig,
		ChainInfo:   chainInfo,
		Path:        path,
		GitConfig:   gitConfig,
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

	r.IconFileName = iconFileName
	r.LogoFileName = logoFileName

	return r
}

func (r *Repo) Commit() *Repo {
	gitcmd.Commit(r.ChainInfo, r.Repo, r.IconFileName, r.LogoFileName, r.GitConfig.Username, r.GitConfig.Email)
	return r
}

func (r *Repo) Push() *Repo {
	gitcmd.Push(r.Repo, r.GitConfig.AccessToken)
	return r
}

func (r *Repo) PullRequest() *Repo {
	err := gitcmd.PullRequest(r.ChainInfo, r.GitConfig.PrTargetBranch, r.GitConfig.AccessToken, r.GitConfig.RepoOrga, r.GitConfig.RepoName)
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
