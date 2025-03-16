package users

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username" gorm:"column:UserName"`
	Email    string `json:"email"`
}

type UserRepository interface {
	GetAllUsers() ([]User, error)
}
