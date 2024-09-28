package models

import "time"

type User struct {
	Id                  int       `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name                string    `gorm:"varchar(100);not null" json:"name"`
	Email               string    `gorm:"varchar(100);unique;not null" json:"email"`
	PasswordHash        string    `gorm:"varchar(450);not null" json:"password_hash"`
	AvatarURL           string    `gorm:"varchar(450)" json:"avatar_url"`
	GoogleCalendarToken string    `gorm:"varchar(450)" json:"google_calendar_token"`
	CreatedAt           time.Time `gorm:"CreatedAt" `
	UpdatedAt           time.Time `gorm:"UpdatedAt" `
	JWTToken            string    `gorm:"varchar(255)"`
}
