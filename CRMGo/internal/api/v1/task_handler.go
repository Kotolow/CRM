package v1

import (
	"CRMGo/internal/models"
	"CRMGo/pkg/database"
	"CRMGo/pkg/mail"
	"CRMGo/pkg/response"
	"CRMGo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	taskService database.TaskRepo
}

func NewTaskHandler(service database.TaskRepo) *TaskHandler {
	return &TaskHandler{taskService: service}
}

// Create godoc
// @Summary      Create task endpoint
// @Description  Creates task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Param        task   body   models.Task true  "Task"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code}/tasks [post]
func (h *TaskHandler) Create(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	err = h.taskService.Create(task, c.Param("code"))
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
// @Summary      Update task endpoint
// @Description  Update task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Param        taskId   path   string true  "Task ID"
// @Param        updatedInfo   body   models.Task true  "Task"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code}/tasks/{taskId} [put]
func (h *TaskHandler) Update(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		utils.InternalError(c, err)
		return
	}
	taskId := c.Param("taskId")
	task.TaskId = taskId

	beforeUpdate, err := h.taskService.FindById(taskId)
	if err != nil {
		if err.Error() == "no task in db" {
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

	err = h.taskService.Update(task)
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	go mail.FormatUpdateMessage(h.taskService, beforeUpdate, c.Param("code"), taskId)

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	c.JSON(http.StatusOK, webResponse)
}

// Delete godoc
// @Summary      Delete task endpoint
// @Description  Delete task
// @Tags         tasks
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Param        taskId   path   string true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code}/tasks/{taskId} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	taskId := c.Param("taskId")
	err := h.taskService.Delete(taskId)
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
// @Summary      Find task by id endpoint
// @Description  Find task by id
// @Tags         tasks
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Param        taskId   path   string true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure		 404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code}/tasks/{taskId} [get]
func (h *TaskHandler) FindById(c *gin.Context) {
	taskId := c.Param("taskId")
	taskResponse, err := h.taskService.FindById(taskId)
	if err != nil {
		if err.Error() == "no task in db" {
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
		Data:   taskResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// FindAll godoc
// @Summary      Find all tasks endpoint
// @Description  Find all tasks
// @Tags         tasks
// @Produce      json
// @Param        code   path   string true  "Project code"
// @Success      200  {object}  response.Response
// @Failure		 404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /v1/projects/{code}/tasks [get]
func (h *TaskHandler) FindAll(c *gin.Context) {
	taskResponse, err := h.taskService.FindAll()
	if err != nil {
		utils.InternalError(c, err)
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: http.StatusText(http.StatusOK),
		Data:   taskResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
