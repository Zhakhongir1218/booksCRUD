package user

import (
	"booksCRUD/pkg/util"
	"fmt"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"os/user"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   State  `json:"status"`
}

type State int8

const (
	UNREGISTERED State = 1
	REGISTERED   State = 2
)

type Dto struct {
	Name string `json:"name"`
}

const hashCost = 10

func SignUp(userName string, password string) Dto {
	userByName := FindUserByName(userName)
	if strings.EqualFold(userByName.Name, userName) {
		_ = fmt.Errorf("user already exists")
	}
	if userByName.Status == REGISTERED {
		_ = fmt.Errorf("user has already registered")
	}
	p := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(p, hashCost)
	var newUser User = User{
		Name:     userName,
		Password: string(hashedPassword),
		Status:   REGISTERED,
	}
	Insert(newUser)
	userDto := Dto{Name: userByName.Name}
	return userDto
}

func SignIn(userName string, password string) string {
	userByName := FindUserByName(userName)
	if userByName.Name == "" {
		panic("System didn't find a user")
	}
	if userByName.Status == UNREGISTERED {
		panic("User was not registered yet")
	}
	p := []byte(password)
	up := []byte(userByName.Password)
	err := bcrypt.CompareHashAndPassword(up, p)
	if err != nil {
		panic(err)
	}
	token, _ := util.CreateToken(userByName.Name, user.User{})
	return token
}

func ReadToken(token string) User {
	decodedToken, _ := util.GetClaimsFromToken(token)
	var name = cast.ToString(decodedToken["name"])
	var password = cast.ToString(decodedToken["password"])
	var status = cast.ToString(decodedToken["status"])

	var s State

	if status == "UNREGISTERED" {
		s = UNREGISTERED
	}
	if status == "REGISTERED" {
		s = REGISTERED
	}

	u := User{
		Name:     name,
		Password: password,
		Status:   s,
	}
	return u
}
