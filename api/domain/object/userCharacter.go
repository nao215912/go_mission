package object

type UserCharacter struct {
	UserID      string `gorm:"column:user_id"`
	CharacterID string `gorm:"column:character_id"`
}

type UserCharacterResponse struct {
	UserCharacterID string `gorm:"column:user_id"`
	CharacterID     string `gorm:"column:character_id"`
	Name            string `gorm:"column:name"`
}
