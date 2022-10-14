package github

type GitConfig struct {
	// Repo
	CloneBranch    string
	PrTargetBranch string
	RepoURL        string
	RepoOrga       string
	RepoName       string

	// User
	Username    string
	Email       string
	AccessToken string
}
