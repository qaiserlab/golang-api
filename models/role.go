package models

type Role struct {
	ID   int    `gorm:"primaryKey;autoIncrement;"`
	Name string `json:"name"`
}
