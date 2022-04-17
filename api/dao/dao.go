package dao

import (
	"go_mission/api/domain/repository"
	"gorm.io/gorm"
)

type Dao interface {
	User() repository.User
	Gacha() repository.Gacha
}

type dao struct {
	db *gorm.DB
}

func New() (Dao, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	return &dao{db: db}, nil
}

func (d *dao) User() repository.User {
	return NewUser(d.db)
}

func (d *dao) Gacha() repository.Gacha {
	return NewGacha(d.db)
}
