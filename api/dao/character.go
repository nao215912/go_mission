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

func (r *character) Create(ctx context.Context, characters []*object.Character) ([]*object.Character, error) {
	tx := r.db.WithContext(ctx).Select("id", "name").Create(&characters)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return characters, nil
}

func (r *character) FindByRand(ctx context.Context, n int) ([]*object.Character, error) {
	const query = `SELECT * FROM characters ORDER BY RAND() LIMIT ?`
	entities := make([]*object.Character, 0, n)
	tx := r.db.WithContext(ctx).Raw(query, n).Scan(&entities)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entities, nil
}
