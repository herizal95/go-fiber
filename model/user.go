package model

import "github.com/google/uuid"

type User struct {
	Uid      uuid.UUID `json:"uid" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"not null"`
	Password string    `json:"password" gorm:"not null"`
}
