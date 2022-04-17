package app

import (
	"go_mission/api/dao"
)

type App struct {
	Dao dao.Dao
}

func NewApp() (*App, error) {
	dao, err := dao.New()
	if err != nil {
		return nil, err
	}

	return &App{Dao: dao}, nil
}
