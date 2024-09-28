package v1

import (
	"CRMGo/pkg/git"
	"CRMGo/pkg/response"
	"CRMGo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GitHubHandler struct {
	gitService git.GitHubRepo
}

func NewGitHubHandler(service git.GitHubRepo) *GitHubHandler {
	return &GitHubHandler{
		gitService: service,
	}
}

// CreateBranch godoc
// @Summary      Create branch
// @Description  Creates branch for the task
// @Tags         git
// @Accept       json
// @Produce      json
// @Param        owner   path   string true  "Owner"
// @Param        repo   path   string true  "Repository"
// @Param        gitInfo   body   git.GitHubInfo true  "Info For GitHub"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/git/{owner}/{repo}/new_branch [post]
func (h *GitHubHandler) CreateBranch(c *gin.Context) {
	var info git.GitHubInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	owner := c.Param("owner")
	repo := c.Param("repo")

	err = h.gitService.CreateBranch(owner, repo, info.SourceBranch, info.BranchName)
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	c.JSON(http.StatusOK, webResponse)
}

// CreatePullRequest godoc
// @Summary      Create PR
// @Description  Creates PR for the task
// @Tags         git
// @Accept       json
// @Produce      json
// @Param        owner   path   string true  "Owner"
// @Param        repo   path   string true  "Repository"
// @Param        gitInfo   body   git.GitHubInfo true  "Info For GitHub"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/git/{owner}/{repo}/new_pr [post]
func (h *GitHubHandler) CreatePullRequest(c *gin.Context) {
	var info git.GitHubInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	owner := c.Param("owner")
	repo := c.Param("repo")

	err = h.gitService.CreatePullRequest(owner, repo, info.SourceBranch, info.TargetBranch, info.Title, info.Body)
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	c.JSON(http.StatusOK, webResponse)
}

// GetRepos godoc
// @Summary      Get Repos
// @Description  Get all repositories connected with organization
// @Tags         git
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/git/repos [get]
func (h *GitHubHandler) GetRepos(c *gin.Context) {
	gitResponse, err := h.gitService.GetRepos()
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   gitResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// GetBranches godoc
// @Summary      Get Branches
// @Description  Get all branches connected with the task
// @Tags         git
// @Produce      json
// @Param        taskId   path   string true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/git/{taskId}/branches [get]
func (h *GitHubHandler) GetBranches(c *gin.Context) {
	taskId := c.Param("taskId")

	repos, err := h.gitService.GetRepos()
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	var gitResponse []git.Branch
	for _, repo := range repos {
		branches, err := h.gitService.GetBranches(repo.Owner, repo.Name, taskId)
		if err != nil {
			utils.InternalError(c, err)
			return
		}
		if len(branches) != 0 {
			gitResponse = append(gitResponse, branches...)
		}
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   gitResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// GetCommits godoc
// @Summary      Get Commits
// @Description  Get all commits connected with the task
// @Tags         git
// @Produce      json
// @Param        taskId   path   string true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/git/{taskId}/commits [get]
func (h *GitHubHandler) GetCommits(c *gin.Context) {
	taskId := c.Param("taskId")

	repos, err := h.gitService.GetRepos()
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	var gitBranches []git.Branch
	for _, repo := range repos {
		branches, err := h.gitService.GetBranches(repo.Owner, repo.Name, taskId)
		if err != nil {
			utils.InternalError(c, err)
			return
		}
		if len(branches) != 0 {
			gitBranches = append(gitBranches, branches...)
		}
	}

	var gitResponse []git.Commit
	for _, branch := range gitBranches {
		commits, err := h.gitService.GetCommits(branch.Owner, branch.Repo, branch.Name, taskId)
		if err != nil {
			utils.InternalError(c, err)
			return
		}
		if len(commits) > 0 {
			gitResponse = append(gitResponse, commits...)
		}
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   gitResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// GetPullRequests godoc
// @Summary      Get PRs
// @Description  Get all Pull Requests connected with the task
// @Tags         git
// @Produce      json
// @Param        taskId   path   string true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/git/{taskId}/pull_requests [get]
func (h *GitHubHandler) GetPullRequests(c *gin.Context) {
	taskId := c.Param("taskId")

	repos, err := h.gitService.GetRepos()
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	var gitResponse []git.PullRequest
	for _, repo := range repos {
		prs, err := h.gitService.GetPullRequests(repo.Owner, repo.Name, taskId)
		if err != nil {
			utils.InternalError(c, err)
			return
		}
		if len(prs) != 0 {
			gitResponse = append(gitResponse, prs...)
		}
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   gitResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
