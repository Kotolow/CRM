package git

type Commit struct {
	Url     string `json:"commit_url"`
	Message string `json:"commit_message"`
	Author  Author `json:"commit_author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PullRequest struct {
	Number    int    `json:"number"`
	User      User   `json:"user"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	State     string `json:"state"`
	Reviewers []User `json:"reviewers"`
}

type User struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	AvatarUrl string `json:"avatar_url"`
}

type GitHubInfo struct {
	SourceBranch string `json:"source_branch,omitempty"`
	TargetBranch string `json:"target_branch,omitempty"`
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	BranchName   string `json:"branch_name,omitempty"`
	TaskId       string `json:"task_id,omitempty"`
}

type Branch struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Repo  string `json:"repo,omitempty"`
	Owner string `json:"owner,omitempty"`
}

type Repository struct {
	Url   string `json:"url"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
