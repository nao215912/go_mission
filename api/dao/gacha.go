package dao

import (
	"go_mission/api/domain/repository"
	"gorm.io/gorm"
)

type gacha struct {
	db *gorm.DB
}

func NewGacha(db *gorm.DB) repository.Gacha {
	return &gacha{db: db}
}
