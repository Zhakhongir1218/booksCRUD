package user

import (
	"booksCRUD/internal/configuration"
	"log"
)

var db = configuration.InitialConnectionPostgres

const insertStatement string = "insert into account (name, password) values ($1, $2)"
const findByUserNameStatement string = "select name, password from account where name=$1"
const updateStatement string = "update account set password = $1 where name=$2"

func Insert(user User) {
	_, err := db.Exec(insertStatement, user.Name, user.Password)
	if err != nil {
		panic(err)
	}
	log.Println("Account was created")
}

func FindUserByName(name string) User {
	row := db.QueryRow(findByUserNameStatement, name)
	u := User{}
	err := row.Scan(&u.Name, &u.Password)
	if err != nil {
		log.Println(err)
		return u
	}
	return u
}

func UpdateUser(user User) {
	_, err := db.Exec(updateStatement, user.Password, user.Name)
	if err != nil {
		panic(err)
	}
}
