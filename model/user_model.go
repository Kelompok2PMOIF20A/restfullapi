package model

type User struct {
	Id       uint   `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Address  string `db:"address" json:"address"`
	Password string `db:"password" json:"password"`
}

type Users []User

type Login struct {
	User  User
	Token string `json:"token"`
}

func (b *User) TableName() string {
	return "Users"
}
