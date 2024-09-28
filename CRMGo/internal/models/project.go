package models

import "time"

type Project struct {
	Id          int    `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"varchar(100);not null"`
	Code        string `gorm:"varchar(100);not null,unique"`
	Description string `gorm:"varchar(500)"`
	CreatedBy   int    `gorm:"not null" json:"created_by,omitempty"`
	user        User   `gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
