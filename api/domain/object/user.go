package object

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string    `gorm:"column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`

	Name  string `gorm:"column:name" json:"name"`
	Token string `gorm:"column:token" json:"token"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	ui, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = ui.String()
	ui, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	u.Token = ui.String()
	return nil
}
