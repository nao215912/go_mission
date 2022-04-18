package repository

import (
	"context"
	"go_mission/api/domain/object"
)

type Character interface {
	Create(ctx context.Context, characters []*object.Character) ([]*object.Character, error)
}
