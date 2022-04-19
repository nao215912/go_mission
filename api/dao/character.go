package dao

import (
	"context"
	"go_mission/api/domain/object"
	"go_mission/api/domain/repository"
	"gorm.io/gorm"
)

type character struct {
	db *gorm.DB
}

func NewCharacter(db *gorm.DB) repository.Character {
	return &character{db: db}
}

// Create Todo:同じアドレスを指しているポインターの是非
func (r *character) Create(ctx context.Context, characters []*object.Character) ([]*object.Character, error) {
	entities := characters
	tx := r.db.WithContext(ctx).Select("name").Create(&entities)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entities, nil
}
