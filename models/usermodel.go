package models

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
