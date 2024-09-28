package database

import (
	"CRMGo/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
)

type ProjectRepo struct {
	Db *gorm.DB
}

func NewProjectRepo(Db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{Db: Db}
}

func (t *ProjectRepo) Create(project models.Project) error {
	result := t.Db.Create(&project)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_project: %v", result.Error.Error()))
		return result.Error
	}
	return nil
}

func (t *ProjectRepo) Update(project models.Project) error {
	result := t.Db.Where("code = ?", project.Code).Updates(project)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_project: %v", result.Error.Error()))
		return result.Error
	}
	return nil
}

func (t *ProjectRepo) Delete(id string) error {
	var project models.Project
	result := t.Db.Where("code = ?", id).Delete(&project)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("db_project: %v", result.Error.Error()))
		return result.Error
	}
	return nil
}

func (t *ProjectRepo) FindById(id string) (models.Project, error) {
	var project models.Project
	result := t.Db.Where("code = ?", id).Last(&project)
	switch {
	case result.RowsAffected == 0:
		return project, errors.New("no project in db")
	case result.Error != nil:
		return project, result.Error
	default:
		return project, nil
	}
}

func (t *ProjectRepo) FindAll() ([]models.Project, error) {
	var projects []models.Project
	results := t.Db.Find(&projects)
	if results.Error != nil {
		slog.Error(fmt.Sprintf("db_project: %v", results.Error.Error()))
		return []models.Project{}, results.Error
	}
	return projects, nil
}
