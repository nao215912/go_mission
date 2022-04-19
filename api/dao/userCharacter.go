package dao

import (
	"context"
	"go_mission/api/domain/object"
	"go_mission/api/domain/repository"
	"gorm.io/gorm"
)

type userCharacter struct {
	db *gorm.DB
}

func NewUserCharacter(db *gorm.DB) repository.UserCharacter {
	return &userCharacter{db: db}
}

// Create Todo:同じアドレスを指しているポインターの是非
func (r *userCharacter) Create(ctx context.Context, userCharacters []*object.UserCharacter) ([]*object.UserCharacter, error) {
	entities := userCharacters
	tx := r.db.WithContext(ctx).Select("user_id", "character_id").Create(&entities)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entities, nil
}

func (r *userCharacter) FindByUsername(ctx context.Context, username string) ([]*object.UserCharacterResponse, error) {
	var entities []*object.UserCharacterResponse
	const query = `
					select
						characters.name,
						user_characters.user_id,
						user_characters.character_id
					from
						user_characters
					join
						characters on user_characters.character_id = characters.id
					join
						users on user_characters.user_id = users.id
					where
						users.name = ?;
					`
	tx := r.db.WithContext(ctx).Raw(query, username).Scan(&entities)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entities, nil
}
