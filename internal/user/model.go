package user

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Dto struct {
	Name string `json:"name"`
}

func SignUp(userName string, password string) Dto {
	user := FindUserByName(userName)
	if strings.EqualFold(user.Name, userName) {
		panic("user already exists")
	}
	p := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(p, 10)
	var newUser User = User{
		Name:     userName,
		Password: string(hashedPassword),
	}
	Insert(newUser)
	userDto := Dto{Name: user.Name}
	return userDto
}
