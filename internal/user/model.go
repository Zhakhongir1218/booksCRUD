package user

import (
	"golang.org/x/crypto/bcrypt"
	"os/user"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Dto struct {
	Name string `json:"name"`
}

const hashCost = 10

func SignUp(userName string, password string) Dto {
	userByName := FindUserByName(userName)
	if strings.EqualFold(userByName.Name, userName) {
		panic("userByName already exists")
	}
	p := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(p, hashCost)
	var newUser User = User{
		Name:     userName,
		Password: string(hashedPassword),
	}
	Insert(newUser)
	userDto := Dto{Name: userByName.Name}
	return userDto
}

func SignIn(userName string, password string) string {
	p := []byte(password)
	userByName := FindUserByName(userName)
	up := []byte(userByName.Password)
	err := bcrypt.CompareHashAndPassword(up, p)
	if err != nil {
		panic(err)
	}
	token, _ := CreateToken(userByName.Name, user.User{})
	return token
}
