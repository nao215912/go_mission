package object

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Character struct {
	ID        uint      `gorm:"column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at"  json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at"  json:"-"`

	Name string `gorm:"column:name;" json:"name"`
	SID  string `json:"characterID"`
}

func (c *Character) AfterCreate(tx *gorm.DB) (err error) {
	c.SID = strconv.FormatUint(uint64(c.ID), 10)
	return nil
}
