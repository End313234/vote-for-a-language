package models

type Language struct {
	Name  string `gorm:"primaryKey;column:name"`
	Votes int    `gorm:"column:votes"`
}
