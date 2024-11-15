package database

import (
	"CRMGo/internal/models"
	"CRMGo/pkg/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
)

type TaskRepo struct {
	Db *gorm.DB
}

func NewTaskRepo(Db *gorm.DB) *TaskRepo {
	return &TaskRepo{Db: Db}
}

func (t *TaskRepo) Create(task models.Task, code string) error {
	var lastTask models.Task
	result := t.Db.Where("project_id = ?", task.ProjectId).Order("id DESC").Last(&lastTask)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_task: %v", result.Error.Error()))
	}
	var newId string
	var err error
	if lastTask.TaskId != "" {
		newId, err = utils.NewTaskName(lastTask.TaskId)
	} else {
		newId, err = utils.NewTaskName(code + "-0")
	}
	if err != nil {
		slog.Error(fmt.Sprintf("db_task: %v", err.Error()))
		return result.Error
	}
	task.TaskId = newId
	result = t.Db.Create(&task)
	return nil
}

func (t *TaskRepo) Update(task models.Task) error {
	result := t.Db.Where("task_id = ?", task.TaskId).Updates(task)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_task: %v", result.Error.Error()))
		return result.Error
	}
	return nil
}

func (t *TaskRepo) Delete(id string) error {
	var task models.Task
	result := t.Db.Where("task_id = ?", id).Delete(&task)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_task: %v", result.Error.Error()))
		return result.Error
	}
	return nil
}

func (t *TaskRepo) FindById(id string) (models.Task, error) {
	var task models.Task
	result := t.Db.Where("task_id = ?", id).Last(&task)
	switch {
	case result.RowsAffected == 0:
		return task, errors.New("no task in db")
	case result.Error != nil:
		return task, result.Error
	default:
		return task, nil
	}
}

func (t *TaskRepo) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	results := t.Db.Find(&tasks)
	if results.Error != nil {
		slog.Error(fmt.Sprintf("db_task: %v", results.Error.Error()))
		return []models.Task{}, results.Error
	}
	return tasks, nil
}

func (t *TaskRepo) FindUserEmailById(id int) (string, error) {
	var u models.User
	result := t.Db.Where("id = ?", id).Last(&u)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_task: %v", result.Error.Error()))
		return "", result.Error
	}
	return u.Email, nil
}
