package database

import (
	"CRMGo/internal/models"
	"CRMGo/pkg/gantt"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type GanttChartRepo struct {
	Db *gorm.DB
}

func NewGanttChartRepo(Db *gorm.DB) *GanttChartRepo {
	return &GanttChartRepo{Db: Db}
}

func (g *GanttChartRepo) GetGanttDataByProject(projectID string) (gantt.GanttChartData, error) {
	var ganttData gantt.GanttChartData
	var project models.Project
	var tasks []models.Task

	result := g.Db.Where("code = ?", projectID).First(&project)
	if result.Error != nil {
		return ganttData, errors.New("project not found")
	}

	result = g.Db.Where("project_id = ?", project.Id).Order("id DESC").Find(&tasks)
	if result.Error != nil {
		return ganttData, result.Error
	}
	fmt.Println(project.Id)

	ganttData.ProjectCode = projectID
	ganttData.ProjectName = project.Name

	for _, task := range tasks {
		assignedUserEmail, err := g.FindUserEmailById(task.AssignedTo)
		if err != nil {
			assignedUserEmail = "Unassigned"
		}

		ganttData.Tasks = append(ganttData.Tasks, gantt.GanttTask{
			TaskID:     task.TaskId,
			Title:      task.Title,
			Status:     task.Status,
			Priority:   task.Priority,
			AssignedTo: assignedUserEmail,
			StartDate:  task.CreatedAt,
			DueDate:    task.DueDate,
			TimeSpent:  task.TimeSpent,
		})
	}

	return ganttData, nil
}

func (g *GanttChartRepo) FindUserEmailById(userID int) (string, error) {
	var user models.User
	result := g.Db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Email, nil
}
