package models

import "time"

type Task struct {
	id          int       `gorm:"primary_key;auto_increment;not null"`
	TaskId      string    `gorm:"varchar(100);not null"`
	ProjectId   int       `gorm:"not null" json:"project_id,omitempty"`
	project     Project   `gorm:"foreignKey:ProjectId;references:Id"`
	Title       string    `gorm:"varchar(100);not null"`
	Description string    `gorm:"varchar(500)"`
	AssignedTo  int       `gorm:"not null" json:"assigned_to,omitempty"`
	user        User      `gorm:"foreignKey:AssignedTo;references:Id"`
	Status      string    `gorm:"varchar(50);not null"`
	Priority    string    `gorm:"varchar(50);not null"`
	DueDate     time.Time `gorm:"timestamp;default:NULL" json:"due_date"`
	TimeSpent   int       `gorm:"int;not null" json:"time_spent"`
	Comments    []Comment `gorm:"serializer:json" json:"comments"`
	CreatedAt   time.Time `gorm:"CreatedAt" `
	UpdatedAt   time.Time `gorm:"UpdatedAt" `
}

type Comment struct {
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
