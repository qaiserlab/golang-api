package models

type User struct {
	ID          int    `gorm:"primaryKey;autoIncrement;"`
	RoleID      string `json:"roleId"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
