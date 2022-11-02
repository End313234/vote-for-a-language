package models

type Language struct {
	Name    string `gorm:"primaryKey;column:name"`
	Votes   int    `gorm:"column:votes"`
	EmojiId string `gorm:"column:emoji_id"`
}
