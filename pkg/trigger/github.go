package trigger

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v60/github" // Ensure you have correct version
	"golang.org/x/oauth2"
)

// GitHubTrigger handles triggering GitHub Actions workflows.
type GitHubTrigger struct {
	client *github.Client
	owner  string
	repo   string
}

// NewGitHubTrigger creates a new trigger instance.
// It expects the GitHub token to be provided, typically via env var.
func NewGitHubTrigger(token, owner, repo string) (*GitHubTrigger, error) {
	if token == "" {
		return nil, fmt.Errorf("github token cannot be empty")
	}
	if owner == "" {
		return nil, fmt.Errorf("github owner cannot be empty")
	}
	if repo == "" {
		return nil, fmt.Errorf("github repo cannot be empty")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GitHubTrigger{
		client: client,
		owner:  owner,
		repo:   repo,
	}, nil
}

// TriggerTokenListUpdate triggers the 'build.yml' workflow.
func (g *GitHubTrigger) TriggerTokenListUpdate(branch string) error {
	if g.client == nil {
		return fmt.Errorf("github client not initialized")
	}
	ctx := context.Background()
	_, err := g.client.Actions.CreateWorkflowDispatchEventByFileName(
		ctx,
		g.owner,
		g.repo,
		"build.yml",
		github.CreateWorkflowDispatchEventRequest{
			Ref: branch,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to dispatch workflow event: %w", err)
	}
	return nil
}

// Helper function to create trigger from environment variables
func NewGitHubTriggerFromEnv() (*GitHubTrigger, error) {
	token := os.Getenv("GITHUB_TOKEN")
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")

	if token == "" || owner == "" || repo == "" {
		// Return nil, nil if not configured, allowing optional trigger usage
		return nil, nil
	}

	return NewGitHubTrigger(token, owner, repo)
}