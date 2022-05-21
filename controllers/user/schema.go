package user

type UserForm struct {
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
