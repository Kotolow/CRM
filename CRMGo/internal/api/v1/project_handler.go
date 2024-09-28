package v1

import (
	"CRMGo/internal/models"
	"CRMGo/pkg/database"
	"CRMGo/pkg/response"
	"CRMGo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectHandler struct {
	projectService database.ProjectRepo
}

func NewProjectHandler(service database.ProjectRepo) *ProjectHandler {
	return &ProjectHandler{projectService: service}
}

// Create godoc
// @Summary      Create project endpoint
// @Description  Creates project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        project   body   models.Project true  "Project"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var project models.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	err = h.projectService.Create(project)
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

// Update godoc
// @Summary      Update project endpoint
// @Description  Update project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Param        updatedInfo   body   models.Project true  "Project"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	var project models.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	projectCode := c.Param("code")
	project.Code = projectCode
	err = h.projectService.Update(project)
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

// Delete godoc
// @Summary      Delete project endpoint
// @Description  Delete project
// @Tags         projects
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
	projectId := c.Param("code")
	err := h.projectService.Delete(projectId)
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

// FindById godoc
// @Summary      Find project by id endpoint
// @Description  Find project by id
// @Tags         projects
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Success      200  {object}  response.Response
// @Failure		 404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code} [get]
func (h *ProjectHandler) FindById(c *gin.Context) {
	projectId := c.Param("code")
	projectResponse, err := h.projectService.FindById(projectId)
	if err != nil {
		if err.Error() == "no project in db" {
			webResponse := response.Response{
				Code:   404,
				Status: http.StatusText(http.StatusNotFound),
				Data:   nil,
			}
			c.JSON(http.StatusNotFound, webResponse)
		} else {
			utils.InternalError(c, err)
		}
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   projectResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// FindAll godoc
// @Summary      Find all projects endpoint
// @Description  Find all projects
// @Tags         projects
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure		 404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects [get]
func (h *ProjectHandler) FindAll(c *gin.Context) {
	projectResponse, err := h.projectService.FindAll()
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   projectResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
