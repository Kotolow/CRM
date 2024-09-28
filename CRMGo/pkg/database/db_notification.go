package database

import (
	"CRMGo/internal/models"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type Notification struct {
	TaskID  string    `json:"task_id"`
	Title   string    `json:"title"`
	Message string    `json:"message"`
	DueDate time.Time `json:"due_date"`
}

type NotificationService struct {
	Db *gorm.DB
}

func NewNotificationService(Db *gorm.DB) *NotificationService {
	return &NotificationService{Db: Db}
}

func (n *NotificationService) GetNotificationsForUser(userID int) ([]Notification, error) {
	var notifications []Notification
	var tasks []models.Task

	now := time.Now()
	oneDayLater := now.Add(24 * time.Hour)
	oneWeekLater := now.Add(7 * 24 * time.Hour)

	result := n.Db.Where("assigned_to = ?", userID).Find(&tasks)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return notifications, result.Error
	}

	for _, task := range tasks {
		if task.Priority == "blocker" {
			notifications = append(notifications, Notification{
				TaskID:  task.TaskId,
				Title:   task.Title,
				Message: "Эта задача является блокирующей. Она находится в статусе 'on hold'.",
				DueDate: task.DueDate,
			})
		}

		if task.DueDate.Before(oneWeekLater) && task.DueDate.After(now) {
			notifications = append(notifications, Notification{
				TaskID:  task.TaskId,
				Title:   task.Title,
				Message: "Deadline is coming in a week:" + task.DueDate.Format("2006-01-02"),
				DueDate: task.DueDate,
			})
		} else if task.DueDate.Before(oneDayLater) && task.DueDate.After(now) {
			notifications = append(notifications, Notification{
				TaskID:  task.TaskId,
				Title:   task.Title,
				Message: "Deadline is coming in a day:" + task.DueDate.Format("2006-01-02"),
				DueDate: task.DueDate,
			})
		}
	}

	return notifications, nil
}
