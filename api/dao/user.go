package dao

import (
	"context"
	"go_mission/api/domain/object"
	"go_mission/api/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
	return &user{db: db}
}

func (r *user) Create(ctx context.Context, user *object.User) (*object.User, error) {
	entity := *user
	tx := r.db.WithContext(ctx).Select("id", "name", "token").Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}

func (r *user) UpdateByName(ctx context.Context, user *object.User, name string) (*object.User, error) {
	entity := *user
	tx := r.db.WithContext(ctx).Model(&entity).Clauses(clause.Returning{}).Where("name = ? and token = ?", user.Name, user.Token).Update("name", name)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}

func (r *user) FindByName(ctx context.Context, name string) (*object.User, error) {
	entity := new(object.User)
	tx := r.db.WithContext(ctx).Where("name = ?", name).Take(entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entity, nil
}

func (r *user) FindByToken(ctx context.Context, token string) (*object.User, error) {
	entity := new(object.User)
	tx := r.db.WithContext(ctx).Where("token = ?", token).Take(entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entity, nil
}
