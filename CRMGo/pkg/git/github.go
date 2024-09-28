package git

import (
	"context"
	"fmt"
	"github.com/google/go-github/v65/github"
	"os"
	"strings"
)

type GitHubRepo struct {
	Client *github.Client
}

func NewGitHubRepo(token string) *GitHubRepo {
	return &GitHubRepo{
		Client: github.NewClient(nil).WithAuthToken(token),
	}
}

func (g *GitHubRepo) CreateBranch(owner, repo, sourceBranch, branchName string) error {
	ctx := context.Background()

	ref, _, err := g.Client.Git.GetRef(ctx, owner, repo, "refs/heads/"+sourceBranch)
	if err != nil {
		return err
	}

	newRef := &github.Reference{
		Ref:    github.String("refs/heads/" + branchName),
		Object: &github.GitObject{SHA: ref.Object.SHA},
	}

	_, _, err = g.Client.Git.CreateRef(ctx, owner, repo, newRef)
	if err != nil {
		return err
	}

	return nil
}

func (g *GitHubRepo) CreatePullRequest(owner, repo, sourceBranch, targetBranch, title, body string) error {
	ctx := context.Background()

	newPR := &github.NewPullRequest{
		Title: github.String(title),
		Head:  github.String(sourceBranch),
		Base:  github.String(targetBranch),
		Body:  github.String(body),
	}

	_, _, err := g.Client.PullRequests.Create(ctx, owner, repo, newPR)
	if err != nil {
		return err
	}

	return nil
}

func (g *GitHubRepo) GetRepos() ([]Repository, error) {
	ctx := context.Background()

	repositories, _, err := g.Client.Repositories.ListByUser(ctx, os.Getenv("GIT_NAME"), nil)
	if err != nil {
		return nil, err
	}
	var repos []Repository
	for _, repo := range repositories {
		repos = append(repos, Repository{
			Url:   *repo.HTMLURL,
			Name:  *repo.Name,
			Owner: *repo.Owner.Login,
		})
	}

	return repos, nil
}

func (g *GitHubRepo) GetBranches(owner, repo, taskID string) ([]Branch, error) {
	ctx := context.Background()

	branches, _, err := g.Client.Repositories.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}

	var filteredBranches []Branch
	for _, branch := range branches {
		if branch.Name != nil {
			if strings.HasPrefix(*branch.Name, taskID) {
				branchURL := fmt.Sprintf("https://github.com/%s/%s/tree/%s", owner, repo, *branch.Name)
				filteredBranches = append(filteredBranches, Branch{
					Name:  *branch.Name,
					Url:   branchURL,
					Owner: owner,
					Repo:  repo,
				})
			}
		}
	}

	return filteredBranches, nil
}

func (g *GitHubRepo) GetCommits(owner, repo, branchName, taskID string) ([]Commit, error) {
	ctx := context.Background()

	commits, _, err := g.Client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{
		SHA: branchName,
	})
	if err != nil {
		return nil, err
	}

	var filteredCommits []Commit
	for _, commit := range commits {
		if commit.Commit != nil && commit.Commit.Message != nil {
			if strings.HasSuffix(*commit.Commit.Message, taskID) {
				formatedCommit := Commit{
					Url:     commit.GetHTMLURL(),
					Message: *commit.Commit.Message,
					Author: Author{
						Name:  commit.Commit.GetAuthor().GetName(),
						Email: commit.Commit.GetAuthor().GetEmail(),
					},
				}
				if formatedCommit.Url != "" {
					filteredCommits = append(filteredCommits, formatedCommit)
				}
			}
		}
	}

	return filteredCommits, nil
}

func (g *GitHubRepo) GetPullRequests(owner, repo, taskID string) ([]PullRequest, error) {
	ctx := context.Background()

	pullRequests, _, err := g.Client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{
		State: "open",
	})
	if err != nil {
		return nil, err
	}

	var filteredPRs []PullRequest
	for _, pr := range pullRequests {
		if pr.Title != nil && strings.HasPrefix(*pr.Title, taskID) {
			reviewers := pr.RequestedReviewers
			formatedReviewers := make([]User, len(reviewers))
			for _, reviewer := range reviewers {
				user := User{
					Name:      reviewer.GetName(),
					Url:       reviewer.GetHTMLURL(),
					AvatarUrl: reviewer.GetAvatarURL(),
				}
				formatedReviewers = append(formatedReviewers, user)
			}
			formatedPR := PullRequest{
				Number: pr.GetNumber(),
				User: User{
					Name:      pr.GetUser().GetName(),
					Url:       pr.GetUser().GetHTMLURL(),
					AvatarUrl: pr.GetUser().GetAvatarURL(),
				},
				Title:     pr.GetTitle(),
				Url:       pr.GetHTMLURL(),
				State:     pr.GetState(),
				Reviewers: formatedReviewers,
			}
			filteredPRs = append(filteredPRs, formatedPR)
		}
	}
	return filteredPRs, nil
}
