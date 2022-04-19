package object

import (
	"gorm.io/gorm"
	"strconv"
)

type UserCharacter struct {
	UserID       uint   `gorm:"column:user_id"`
	CharacterID  uint   `gorm:"column:character_id"`
	SUserID      string `json:"UserID"`
	SCharacterID string `json:"characterID"`
}

type UserCharacterResponse struct {
	UserCharacterID string `gorm:"column:user_id"`
	CharacterID     string `gorm:"column:character_id"`
	Name            string `gorm:"column:name"`
}

func (u *UserCharacter) AfterCreate(tx *gorm.DB) (err error) {
	u.SUserID = strconv.FormatUint(uint64(u.UserID), 10)
	u.SCharacterID = strconv.FormatUint(uint64(u.CharacterID), 10)
	return nil
}
