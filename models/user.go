package models

type User struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	RoleID      int
	Role        Role
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
}
