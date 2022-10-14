package cmd

import (
	"context"
	"fmt"

	"github.com/forbole/bdtool/types"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func PullRequest(chain *types.ChainInfo, targetBranch string, AccessToken string, organization string, repoName string) error {
	client := github.NewClient(oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: AccessToken,
			},
		),
	))

	// Create PR
	pr, _, err := client.PullRequests.Create(
		context.Background(),
		organization,
		repoName,
		&github.NewPullRequest{
			Title:               github.String(fmt.Sprintf("Create new chain config: %s-%s.json", chain.Name, chain.Type)),
			Head:                github.String(fmt.Sprintf("refs/heads/chains/%s/%s", chain.Name, chain.Type)),
			Base:                github.String(targetBranch),
			MaintainerCanModify: github.Bool(true),
		},
	)
	if err != nil {
		return fmt.Errorf("error while creating pull request: %s", err)
	}

	// Add label "chain config"
	_, _, err = client.Issues.AddLabelsToIssue(context.Background(), "forbole", "big-dipper-2.0-cosmos", *pr.Number, []string{"chain config"})
	if err != nil {
		return fmt.Errorf("error while adding label(s): %s", err)
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
	return nil
}
