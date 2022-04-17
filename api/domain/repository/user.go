package repository

import (
	"context"
	"go_mission/api/domain/object"
)

type User interface {
	Create(ctx context.Context, user *object.User) (*object.User, error)
	UpdateByName(ctx context.Context, user *object.User, name string) (*object.User, error)
}
