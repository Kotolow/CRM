package mail

import (
	"CRMGo/internal/models"
	"CRMGo/pkg/database"
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
	"strings"
)

func FormatUpdateMessage(db database.TaskRepo, beforeUpdate models.Task, projectCode, taskId string) {
	task, err := db.FindById(taskId)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	userMail, err := db.FindUserEmailById(task.AssignedTo)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	changes := compareTasks(beforeUpdate, task)
	if len(changes) == 0 {
		slog.Info("No changes detected in task.")
		return
	}

	subject := fmt.Sprintf("Update of %s in %s project", taskId, projectCode)
	body := fmt.Sprintf("The following changes were made to task %s in project %s:\n\n%s", taskId, projectCode, changes)

	err = SendEmail(userMail, subject, body)
	if err != nil {
		slog.Error("Failed to send email: ", err)
	} else {
		slog.Info("Email sent successfully.")
	}
}

func compareTasks(before, after models.Task) string {
	var changes []string

	if before.Title != after.Title {
		changes = append(changes, fmt.Sprintf("Title: %s -> %s", before.Title, after.Title))
	}
	if before.Description != after.Description {
		changes = append(changes, fmt.Sprintf("Description: %s -> %s", before.Description, after.Description))
	}
	if before.AssignedTo != after.AssignedTo {
		changes = append(changes, fmt.Sprintf("AssignedTo: %d -> %d", before.AssignedTo, after.AssignedTo))
	}
	if before.Status != after.Status {
		changes = append(changes, fmt.Sprintf("Status: %s -> %s", before.Status, after.Status))
	}
	if before.Priority != after.Priority {
		changes = append(changes, fmt.Sprintf("Priority: %s -> %s", before.Priority, after.Priority))
	}
	if before.DueDate != after.DueDate {
		changes = append(changes, fmt.Sprintf("DueDate: %v -> %v", before.DueDate, after.DueDate))
	}
	if before.TimeSpent != after.TimeSpent {
		changes = append(changes, fmt.Sprintf("TimeSpent: %d -> %d", before.TimeSpent, after.TimeSpent))
	}
	if len(before.Comments) != len(after.Comments) {
		changes = append(changes, fmt.Sprintf("Comments: %v -> %v", before.Comments, after.Comments))
	}

	if len(changes) == 0 {
		return ""
	}

	return strings.Join(changes, "\n")
}

func SendEmail(to string, subject string, body string) error {
	from := "h&h@gmail.com"
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
