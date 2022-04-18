package object

import "time"

type Character struct {
	ID        uint      `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Name        string `gorm:"column:name;"`
	CharacterID string `gorm:"column:characterID"`
}
