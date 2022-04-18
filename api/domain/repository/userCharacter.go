package repository

import (
	"context"
	"go_mission/api/domain/object"
)

type UserCharacter interface {
	Create(ctx context.Context, userCharacters []*object.UserCharacter) ([]*object.UserCharacter, error)
}
