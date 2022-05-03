package object

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Character struct {
	ID        string    `gorm:"column:id" json:"characterID"`
	CreatedAt time.Time `gorm:"column:created_at"  json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at"  json:"-"`

	Name string `gorm:"column:name;" json:"name"`
}

func (c *Character) BeforeCreate(tx *gorm.DB) (err error) {
	ui, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	c.ID = ui.String()
	return nil
}
