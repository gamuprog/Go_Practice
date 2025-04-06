package model

import (
	"Go_Practice/domain/entity"
	"time"
)

type User struct {
	UserID         string `gorm:"primaryKey;size:26"`
	Email          string `gorm:"size:255;unique;not null"`
	HashedPassword string `gorm:"size:255;not null"`
	IsDeleted      bool   `gorm:"default:false;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *User) Entity() entity.User {
	return entity.User{
		UserID:         u.UserID,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		IsDeleted:      u.IsDeleted,
	}
}