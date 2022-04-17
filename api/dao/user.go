package dao

import (
	"go_mission/api/domain/repository"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
	return &user{db: db}
}
