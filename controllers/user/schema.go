package user

type UserForm struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
