package object

type UserCharacter struct {
	UserID      uint `gorm:"column:user_id"`
	CharacterID uint `gorm:"column:character_id"`
}
